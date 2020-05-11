package feed

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

func Crawler(url string, isAtom bool) {
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
	fmt.Println(feedEntries)
}

type Entry struct {
	Title string
	Url   string
	Date  time.Time
}
