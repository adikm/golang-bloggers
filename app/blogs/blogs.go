package blogs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Blogs struct {
	Blog []Blog `json:"blogs"`
}

type Blog struct {
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
	Rss       string `json:"rss"`
}

func Get() Blogs {
	var blogs Blogs
	resp, err := http.Get("https://raw.githubusercontent.com/adikm/golang-bloggers/master/blogs/blogs.json")
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return Blogs{}
	}
	_ = json.NewDecoder(resp.Body).Decode(&blogs)
	return blogs
}
