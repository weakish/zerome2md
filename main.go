package main

import (
	"encoding/json"
	"fmt"
	"github.com/weakish/goaround"
	"github.com/weakish/gosugar"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type post struct {
	PostId    int `json:"post_id"`
	DateAdded int64 `json:"date_added"`
	Body      string `json:"body"`
	Meta      string `json:"meta"`
}

func main() {
	bytes, err := ioutil.ReadFile("data.json")
	goaround.PanicIf(err)

	var user map[string]json.RawMessage
	goaround.PanicIf(json.Unmarshal(bytes, &user))
	var posts []post
	goaround.PanicIf(json.Unmarshal(user["post"], &posts))

	var images []string
	images, err = filepath.Glob("*.jpg")
	goaround.PanicIf(err)

	var imageIds gosugar.StringSet = gosugar.NewStringSet()
	for _, image := range images {
		imageIds.Add(strings.TrimSuffix(image, ".jpg"))
	}

	for _, p := range posts {
		fmt.Println(p.Body + "\n")

		var id string = strconv.Itoa(p.PostId)
		if imageIds.Contains(id) {
			fmt.Println("![attached image](" + id + ".jpg)\n")
		}

		var date time.Time = time.Unix(p.DateAdded, 0).UTC()
		fmt.Println("-- " + date.Format("2006-01-02") + "\n\n")
	}
}
