package feed

import (
	"encoding/xml"
	"fmt"
	"time"
)

func getAtomEntries(decoder xml.Decoder) []Entry {
	feed := AtomFeed{}

	err := decoder.Decode(&feed)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
	}
	var entries []Entry
	for _, item := range feed.Entry {
		date, _ := time.Parse(time.RFC3339, item.PublishedAt)
		entries = append(entries, Entry{item.Title, item.Link.Url, date})
	}

	return entries
}

type Link struct {
	Url string `xml:"href,attr"`
}

type AtomEntry struct {
	Title       string `xml:"title"`
	Link        Link   `xml:"link"`
	PublishedAt string `xml:"updated"`
}

type AtomFeed struct {
	Entry []AtomEntry `xml:"entry"`
}
