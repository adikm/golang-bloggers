package feed

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func GetFeed(url string, isAtom bool, entries chan []Entry) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var feedEntries []Entry
	decoder := xml.NewDecoder(resp.Body)
	if isAtom {
		feedEntries = append(feedEntries, getAtomEntries(*decoder)...)
	} else {
		feedEntries = append(feedEntries, getRssEntries(*decoder)...)
	}
	entries <- filterLastWeekOnly(&feedEntries)
}

func filterLastWeekOnly(entries *[]Entry) []Entry {
	weekAgo := time.Now().Add(-14 * 24 * time.Hour)
	var lastWeekEntries []Entry
	for _, entry := range *entries {
		isWithinLastWeek := entry.Date.After(weekAgo)
		if isWithinLastWeek {
			lastWeekEntries = append(lastWeekEntries, entry)
		}
	}
	return lastWeekEntries
}

type Entry struct {
	Title string
	Url   string
	Date  time.Time
}
