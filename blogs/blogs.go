package blogs

import (
	"encoding/json"
	"io/ioutil"
	"os"
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
	open, _ := os.Open("blogs/blogs.json")
	dat, _ := ioutil.ReadAll(open)

	_ = json.Unmarshal(dat, &blogs)

	defer open.Close()
	return blogs
}
