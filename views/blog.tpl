<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
<head>
	<meta http-equiv="content-type" content="text/html; charset=utf-8" />
	<meta name="keywords" content="" />
	<meta name="description" content="" />
	<link rel="stylesheet" type="text/css" href="<%CUSTOM_URL_CSS%>default.css" />
	<link rel="stylesheet" type="text/css" href="<%CUSTOM_URL_CSS%>blog.css" />
	<link rel="shortcut icon" href="<%CUSTOM_URL_IMG%>i_logo1.png" >
	<title><%.PageTitle%></title>
</head>
<body>
<div id="header">
	<div id="logo">
		<h1><img href="/" src="<%CUSTOM_URL_IMG%>ckeyer.png" alt="ckeyer" /></h1>
		<h2><a href="/">Man, I just luv technology...... </a></h2>
	</div>
	<div id="menu">
		<ul>
			<li class="first"><a href="/blog">My Blog</a></li>
			<li><a href="/photo">My Photos</a></li>
			<li><a href="/favorite">My Favorites</a></li>
			<li><a href="/contact">Contact Me</a></li>
			<li><a href="/about">About Me</a></li>
			<li><a></a></li>
		</ul>
	</div>
</div>
<div id="content">
	<div class="colOne">
		<div class="art_area">
		  	<div class="article">
		  		<div class="art_title">
		  			<a href="#">一个被诅咒的IP</a>
		  		</div>
		  		<div class="art_time">2015年1月11日 22:17:44 &nbsp&nbsp
		  			<span> 
					<% range $index, $elem := .Tags %>
			  			<a href="/tag/<%$elem.ID%>" class="art_label"><%$elem.Name%></a>&nbsp
					<%end%>
		  			</span>
	  			</div>
		  		<div class="art_content">
		  			<% str2html ( .ArticleContent|DECODEBASE64 ) %>
		  		</div>  
			</div>
			<div class="next_art">
		  		<a href="/blog/<%  .Previous.ID%>" title="">上一篇：<%  .Previous.Title|DECODEBASE64  %></a>
			</div>
			<div class="next_art">
		  		<a href="/blog/<%  .Next.ID%>" title="">下一篇：<%  .Next.Title|DECODEBASE64  %></a>
			</div>
		</div>
	</div>
	
	<div id="colTwo">
		<ul>
			<li>
				<h2>Archives</h2>
				<ul>
					<li><a href="#">December 2014</a></li>
					<li><a href="#">November 2014</a></li>
					<li><a href="#">October 2014</a></li>
					<li><a href="#">September 2014</a></li>
					<li><a href="#">August 2014</a></li>
				</ul>
			</li>
			<li>
				<h2>Categories</h2>
				<ul>
					<% range $index, $elem := .HotTags %>
						<li><a href="/tag/<%$elem.ID%>" ><%$elem.Name%></a> (<%$elem.ArtCount%>)</li>
					<%end%>
				</ul>
			</li>

			<li id="Friend_site_Link">
				<h2>Friend Site Link </h2>
				<ul>
					<li><a href="http://ys.cjstudio.org/" target="_blank">严申的个人博客</a></li>
					<li><a href="http://www.originate.com/">Originate</a></li>
					<li><a href="http://programmer.csdn.net/programmer.html">程序员杂志</a></li>
					<li><a href="http://www.litrin.net/">开源小站</a></li>
					<li><a href="http://lusongsong.com/">卢松松博客</a></li>
					<li><a href="https://gowalker.org/">gowalker.org</a></li>
				</ul>
			</li>
		</ul>
		<div style="clear: both;">&nbsp;</div>
	</div>
</div>
<div id="footer"><br>
	<p>Copyright &copy; 2014 lab204. Designed by <a href="#"><strong>Lab204-CJStudio</strong></a></p>
</div>
</body>
</html>