package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	nprof "net/http/pprof"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

const (
	API_PREFIX = "/api"
)

var (
	fh = http.FileServer(http.Dir("./static/"))
)

func NewM() *martini.ClassicMartini {
	r := martini.NewRouter()
	m := martini.New()
	//m.Use(martiniLogger())
	m.Use(martini.Recovery())
	m.Use(render.Renderer())
	m.MapTo(r, (*martini.Routes)(nil))
	m.Action(r.Handle)
	return &martini.ClassicMartini{Martini: m, Router: r}

}

func StaticHandle(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("static")
	name := strings.TrimLeft(req.URL.Path, "/")
	if name == "" || strings.HasSuffix(name, "/") {
		name += "index.html"
	}
	name = "static/" + name
	ext := path.Ext(name)
	ct := mime.TypeByExtension(ext)
	rw.Header().Set("Content-Type", ct)
	data, err := ioutil.ReadFile(name)
	if err != nil {
		rw.WriteHeader(404)
		return
	}
	rw.Write(data)
	return
	hdr := req.Header.Get("Accept-Encoding")
	if strings.Contains(hdr, "gzip") {
		rw.Header().Set("Content-Encoding", "gzip")
		rw.Write(data)
	} else {
		gz, err := gzip.NewReader(bytes.NewBuffer(data))
		if err != nil {
			rw.Write([]byte(err.Error()))
			return
		}
		io.Copy(rw, gz)
		gz.Close()
	}
}

type User struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Created time.Time `json:"created"`
}

func main() {

	m := NewM()
	m.Use(Auth)
	headers := make(map[string]string)
	headers["Access-Control-Allow-Methods"] = "GET,OPTIONS,POST,DELETE"
	headers["Access-Control-Allow-Credentials"] = "true"
	headers["Access-Control-Max-Age"] = "864000"
	headers["Access-Control-Expose-Headers"] = "Record-Count,Limt,Offset,Content-Type"
	headers["Access-Control-Allow-Headers"] = "Limt,Offset,Content-Type,Origin,Accept,Authorization"

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     strings.Split(headers["Access-Control-Allow-Origin"], ","),
		AllowMethods:     strings.Split(headers["Access-Control-Allow-Methods"], ","),
		AllowHeaders:     strings.Split(headers["Access-Control-Allow-Headers"], ","),
		ExposeHeaders:    strings.Split(headers["Access-Control-Expose-Headers"], ","),
		AllowCredentials: true,
		MaxAge:           time.Second * 864000,
	}))

	m.NotFound(StaticHandle)

	m.Group("/api", func(r martini.Router) {
		r.Get("/hello", Hello)
	})
	m.Group("/prof", func(r martini.Router) {
		r.Get("/goroutine", nprof.Handler("goroutine").ServeHTTP)
		r.Get("/heap", nprof.Handler("heap").ServeHTTP)
		r.Get("/threadcreate", nprof.Handler("threadcreate").ServeHTTP)
		r.Get("/block", nprof.Handler("block").ServeHTTP)
		r.Get("/cpu", func(rw http.ResponseWriter, req *http.Request) {
			// pprof.StopCPUProfile()
		})
	})
	runtime.SetBlockProfileRate(1)
	pprof.StartCPUProfile(os.Stdout)
	m.RunOnAddr(":8000")
}

func Hello(c martini.Context, rw http.ResponseWriter, req *http.Request, rnd render.Render) {
	fmt.Println("hello, ", req.URL.String())
	u := &User{
		Name:    "ckeyer",
		Age:     112,
		Created: time.Now(),
	}
	rnd.JSON(200, u)
}

func Auth(c martini.Context, rw http.ResponseWriter, req *http.Request) {
	fmt.Println("auth handle: ", req.URL.String())
}

func Login(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("login...")
	rw.Write([]byte("login...\n"))
}
