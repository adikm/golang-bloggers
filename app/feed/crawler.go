package feed

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func GetFeed(url string, isAtom bool) []Entry {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	var feedEntries []Entry
	decoder := xml.NewDecoder(resp.Body)
	if isAtom {
		feedEntries = append(feedEntries, getAtomEntries(*decoder)...)
	} else {
		feedEntries = append(feedEntries, getRssEntries(*decoder)...)
	}
	return feedEntries
}

type Entry struct {
	Title string
	Url   string
	Date  time.Time
}
