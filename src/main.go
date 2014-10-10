package GroupMengine

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type HandleFunc func(client *http.Client, term string, w http.ResponseWriter) string

type GroupMengineConfig struct {
	Handlers map[string]HandleFunc
	Bots     map[string]string
}

var Config GroupMengineConfig

func init() {
	Config = GetConfig()
	http.HandleFunc("/newmsg", sendMessage)
}

type NewMessage struct {
	Id          string `json:"id"`
	Source_guid string `json:"source_guid"`
	Created_at  int    `json:"created_at"`
	User_id     string `json:"user_id"`
	Group_id    string `json:"group_id"`
	Name        string `json:"name"`
	Avatar_url  string `json:"avatar_url"`
	Text        string `json:"text"`
}

func sendMessage(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	p := make([]byte, r.ContentLength)
	_, err := r.Body.Read(p)
	if err == nil {
		var msg NewMessage
		err1 := json.Unmarshal(p, &msg)
		if err1 == nil {
			// Save off message :)
			_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Messages", nil), &msg)
			if err != nil {
				return
			}
			cmd := msg.Text
			if strings.Index(cmd, "/") == 0 {
				cmdType := strings.Split(cmd, " ")[0]
				cmdBody := strings.TrimSpace(strings.Replace(cmd, cmdType, "", -1))
				cmdBody = strings.Replace(cmdBody, " ", "+", -1)
				handler, ok := Config.Handlers[cmdType]
				if ok {
					form := make(url.Values)
					form.Add("bot_id", Config.Bots[msg.Group_id])
					form.Add("text", handler(client, cmdBody, w))
					client.PostForm("https://api.groupme.com/v3/bots/post", form)
				}
			}
		}
	}
	r.Body.Close()
}
