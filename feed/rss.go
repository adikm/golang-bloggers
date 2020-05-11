package feed

import (
	"encoding/xml"
	"fmt"
	"time"
)

func getRssEntries(decoder xml.Decoder) []Entry {
	feed := Rss{}

	err := decoder.Decode(&feed)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
	}
	var entries []Entry
	for _, item := range feed.Channel.Items {
		date, _ := time.Parse(time.RFC1123Z, item.PublishedAt)
		entries = append(entries, Entry{item.Title, item.Link, date})
	}

	return entries
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishedAt string `xml:"pubDate"`
}

type RssChannel struct {
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel RssChannel `xml:"channel"`
}
