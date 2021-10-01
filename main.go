package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/weakish/gosugar"
)

type post struct {
	PostId    int    `json:"post_id"`
	DateAdded int64  `json:"date_added"`
	Body      string `json:"body"`
	Meta      string `json:"meta"`
}

func main() {
	bytes, _ := ioutil.ReadFile("data.json")

	var user map[string]json.RawMessage
	_ = json.Unmarshal(bytes, &user)
	var posts []post
	_ = json.Unmarshal(user["post"], &posts)

	var images []string
	images, _ = filepath.Glob("*.jpg")

	var imageIds gosugar.StringSet = gosugar.NewStringSet()
	for _, image := range images {
		imageIds.Add(strings.TrimSuffix(image, ".jpg"))
	}

	prefix := flag.String("p", "", "site prefix")
	repo := flag.String("r", ".", "repository directory")
	name := flag.String("n", "zerome", "git author name")
	email := flag.String("e", "", "git author email")
	flag.Parse()
	for _, p := range posts {
		var id string = strconv.Itoa(p.PostId)
		var date time.Time = time.Unix(p.DateAdded, 0).UTC()
		if *prefix == "" {
			fmt.Println(p.Body + "\n")
			if imageIds.Contains(id) {
				fmt.Println("![attached image](" + id + ".jpg)\n")
			}
			fmt.Println("-- " + date.Format("2006-01-02") + "\n\n")
		} else {
			body := p.Body + "\n"
			if imageIds.Contains(id) {
				body += "![attached image](" + *prefix + "/" + id + ".jpg)\n"
			}

			r, _ := git.PlainOpen(*repo)
			w, _ := r.Worktree()
			w.Commit(body, &git.CommitOptions{
				Author: &object.Signature{
					Name:  *name,
					Email: *email,
					When:  date,
				},
			})
		}
	}
}
