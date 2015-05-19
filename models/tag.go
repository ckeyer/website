/*
**/

package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Tag struct {
	Id       int64
	ParentId int    `orm:"default(0)"`
	Name     string `orm:"index;size(32);unique"`

	// Blogs []*Blog `orm:"reverse(many)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewTag(name string, parent int) *Tag {
	return &Tag{Name: name, ParentId: parent}
}
func GetTagById(id int64) (tag *Tag, err error) {
	o := orm.NewOrm()
	tag = &Tag{Id: id}
	err = o.Read(tag)
	return
}
func GetTagIdByName(name string) int64 {
	tag := &Tag{}
	o := orm.NewOrm()

	qs := o.QueryTable("tag").Filter("name", strings.ToLower(name))
	err := qs.One(tag)
	// err := o.Read(tag)
	if err != nil {
		// log.Println(err.Error())
		tag := &Tag{Name: strings.ToLower(name)}
		tag.Insert()
	}
	// log.Println("add log in there", tag)
	return tag.Id
}
func (this *Tag) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err != nil {
		log.Println(err.Error())
	}
	if err == nil {
		this.Id = id
	}
	return err
}
func (this *Tag) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}

func (this *Tag) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
func GetAllBlogsById(blog_id int) []*Blog {
	blogs := make([]*Blog)
	return
}

func (this *Tag) GetAllBlogsByName(name string) {

}
