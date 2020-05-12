package api

import (
	"encoding/json"
	"github.com/adikm/golang-bloggers/app/blogs"
	"github.com/adikm/golang-bloggers/app/feed"
	"net/http"
	"strings"
)

func InitServer() {
	http.HandleFunc("/feed", feedHandler)
	_ = http.ListenAndServe(":3000", nil)
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	blogsList := blogs.Get()

	var feedEntries []feed.Entry
	for _, blog := range blogsList.Blog {
		feedEntries = append(feedEntries, feed.GetFeed(blog.Rss, strings.HasSuffix(blog.Rss, "atom"))...)
	}
	response, _ := json.Marshal(feedEntries)
	w.Write(response)
}
