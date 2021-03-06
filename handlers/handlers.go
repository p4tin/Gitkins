package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/p4tin/Gitkins/clients"
	"github.com/p4tin/Gitkins/config"
	"io"
)

type HealthInfo struct {
	Version     string       `json:"version,omitempty"`
	JenkinsUrl  string       `json:"jenkinsUrl,omitempty"`
	JenkinsUser string       `json:"jenkinsUser,omitempty"`
	Watches     []config.Job `json:"watches,omitempty"`
}

func init() {

}

func HealthEventHandler(w http.ResponseWriter, r *http.Request) {
	info := HealthInfo{
		Version: config.Version,
		Watches: config.Config.Watches,
	}
	b, err := json.MarshalIndent(info, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	io.WriteString(w, string(b))
}

/*

	Github Event Headers:
		- X-GitHub-Event == PullRequestEvent
		- X-GitHub-Delivery == GUID of event
		- X-Hub-Signature == HMAC of the event body

*/
func GitEventHandler(w http.ResponseWriter, r *http.Request) {
	event_type := r.Header.Get("X-GitHub-Event")

	if event_type == "pull_request" {
		pr_event := new(github.PullRequestEvent)
		json.NewDecoder(r.Body).Decode(pr_event)

		if config.Config.Debug {
			log.Printf("Event Type: %s, Created by: %s\n", event_type, pr_event.PullRequest.Base.User.Login)
			log.Printf("Message: %s\n", r.Body)
		}

		go clients.ProcessPullRequest(*pr_event)
		log.Println("Handler exiting...")
	} else {
		log.Printf("Event %s not supported yet.\n", event_type)
	}
}
