package GroupMengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type QuoteMessage struct {
	Quote string
}

func getQuote(client *http.Client, w http.ResponseWriter) string {
	searchUrl := "http://www.iheartquotes.com/api/v1/random.json?source=south_park+oneliners+riddles"
	res, err := client.Get(searchUrl)
	if err != nil {
		fmt.Fprint(w, "get failed")
	}

	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprint(w, "get failed")
	}

	var quoteMsg QuoteMessage
	err = json.Unmarshal(resp, &quoteMsg)
	return quoteMsg.Quote
}

func randoText(client *http.Client, term string, w http.ResponseWriter) string {
	
	rando := getQuote(client, w)
	snarks := [...]string{"I LOVE YOU",
		"How dare you",
		"You better have that /up ready",
		"/say dingus",
		"May the odds ever...",
		"v1.1.2-r12",
		rando,
	}
	rand.Seed(time.Now().Unix())
	return snarks[rand.Intn(len(snarks))]
}