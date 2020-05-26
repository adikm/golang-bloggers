package api

import (
	"encoding/json"
	"github.com/adikm/golang-bloggers/app/blogs"
	"github.com/adikm/golang-bloggers/app/feed"
	"github.com/adikm/golang-bloggers/db"
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

	feedEntriesChannel := make(chan []feed.Entry, len(blogsList.Blog))
	var feedEntries []feed.Entry
	for _, blog := range blogsList.Blog {
		go feed.GetFeed(blog.Rss, strings.HasSuffix(blog.Rss, "atom"), feedEntriesChannel)
		feedEntries = append(feedEntries, <-feedEntriesChannel...)
	}
	db.InsertEntries(&feedEntries)
	response, _ := json.Marshal(feedEntries)
	w.Write(response)
}
