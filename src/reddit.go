// reddit.go
package GroupMengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type RedditResponse struct {
	Data RedditData
}

type RedditData struct {
	Children []RedditChildren
}

type RedditChildren struct {
	Data RedditChildData
}

type RedditChildData struct {
	Body string
}

func redditSearch(client *http.Client, term string, w http.ResponseWriter) string {
	searchUrl := fmt.Sprintf("http://www.reddit.com/r/all/comments.json?sort=random")
	res, err := client.Get(searchUrl)
	if err != nil {
		fmt.Fprint(w, "get failed")
	}

	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprint(w, "bad resp body")
	}

	var redditResp RedditResponse
	err = json.Unmarshal(resp, &redditResp)

	rand.Seed(time.Now().Unix())
	numComments := len(redditResp.Data.Children)
	if numComments > 0 {
		randPos := rand.Intn(numComments)

		return redditResp.Data.Children[randPos].Data.Body
	} else {
		return ""
	}
}
