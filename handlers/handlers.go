package handlers

import (
	"net/http"
	"github.com/yosida95/golang-jenkins"
	"fmt"
	"strconv"
	"time"
	"log"
	"encoding/json"

	"github.com/google/go-github/github"
)

func GitEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Message: %+v\n", r.Body)
	v := new(github.PullRequest)
	json.NewDecoder(r.Body).Decode(v)

	if v.Number != nil {
		pr, _, err := client.PullRequests.Get("pfortin-urbn", "hooktesting", *v.Number)
		if err != nil {
			panic(err)
		}
		if *pr.State == "open" {
			log.Println(*pr.State)
			status1 := &github.RepoStatus{State: &pending, TargetURL: &targetUrl, Description: &pendingDesc, Context: &appName}
			client.Repositories.CreateStatus("pfortin-urbn", "hooktesting", *pr.Head.SHA, status1)
			log.Println(pr)

			jauth := &gojenkins.Auth{ApiToken: jenkinsApiToken, Username: jenkinsUsername}
			jenkins := gojenkins.NewJenkins(jauth, jenkinsBaseUrl)
			job, err := jenkins.GetJob("PaulTestJob")
			if err != nil {
				log.Printf("Jenkins ERROR: %+v\n", err)
			} else {
				fmt.Printf("%s\n", job)
				params := url.Values{
					"PR": []string{strconv.Itoa(*v.Number)},
					"SHA": []string{"test1"},
				}
				status := jenkins.Build(job, params)
				log.Printf("Status: %+v\n", status)
			}


			time.Sleep(10 * time.Second)
			status2 := &github.RepoStatus{State: &success, TargetURL: &targetUrl, Description: &successDesc, Context: &appName}
			client.Repositories.CreateStatus("pfortin-urbn", "hooktesting", *pr.Head.SHA, status2)
		} else {
			log.Printf("PR event but not open state. (%s)", *pr.State)
		}
	} else {
		log.Printf("Hook sent event that was not a PR - %+v\n", v)
	}
}
