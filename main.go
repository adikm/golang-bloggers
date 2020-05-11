package main

import (
	"github.com/adikm/golang-bloggers/blogs"
	"github.com/adikm/golang-bloggers/feed"
	"strings"
)

func main() {
	blogaa := blogs.Get()

	for _, blog := range blogaa.Blog {
		feed.Crawler(blog.Rss, strings.HasSuffix(blog.Rss, "atom"))
	}

}
