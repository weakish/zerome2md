package main

import (
	"encoding/json"
	"fmt"
	"github.com/weakish/goaround"
	"io/ioutil"
	"os"
	"time"
)

type post struct {
	PostId    int `json:"post_id"`
	DateAdded int64 `json:"date_added"`
	Body      string `json:"body"`
	Meta      string `json:"meta"`
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	goaround.PanicIf(err)

	var user map[string]json.RawMessage
	goaround.PanicIf(json.Unmarshal(bytes, &user))
	var posts []post
	goaround.PanicIf(json.Unmarshal(user["post"], &posts))

	for _, p := range posts {
		fmt.Println(p.Body + "\n")
		var date time.Time = time.Unix(p.DateAdded, 0).UTC()
		fmt.Println("-- " + date.Format("2006-01-02") + "\n\n")
	}
}
