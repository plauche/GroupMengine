package GroupMengine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type SpotifyUrl struct {
	Spotify string
}

type SpotifyItem struct {
	External_urls SpotifyUrl
}

type SpotifyTracks struct {
	Href  string
	Items []SpotifyItem
}

type SpotifyResponse struct {
	Tracks SpotifyTracks
}

func spotifySearch(client *http.Client, term string, w http.ResponseWriter) string {
	searchUrl := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track", term)
	res, err := client.Get(searchUrl)
	if err != nil {
		fmt.Fprint(w, "get failed")
	}

	resp, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Fprint(w, "bad resp body")
	}

	var spotResponse SpotifyResponse
	err = json.Unmarshal(resp, &spotResponse)

	rand.Seed(time.Now().Unix())
	numTracks := len(spotResponse.Tracks.Items)
	if numTracks > 0 {
		randPos := rand.Intn(numTracks)

		return spotResponse.Tracks.Items[randPos].External_urls.Spotify
	} else {
		return ""
	}
}
