package clients

import (
	"log"
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"

	"github.com/p4tin/Gitkins/config"
)

var (
	pending = "pending"
	success = "success"
	failure = "error"
	targetUrl = "https://urbn-ci.ngrok.com/status"
	pendingDesc = "Build/testing in progress, please wait."
	successDesc = "Build/testing successful."
	failureDesc = "Build or Unit Test failed."
	appName = "Urbn-CI"
)

func init() {

}

func ProcessPullRequest(pr_evt github.PullRequestEvent) {
	//Get the Pull Request
	log.Printf("%s - %s\n",  *pr_evt.PullRequest.Head.User.Login, *pr_evt.PullRequest.Head.Repo.Name)
	watch := -1;
	for i, w := range config.Config.Watches {
		if *pr_evt.PullRequest.Head.User.Login == w.GithubAccount && *pr_evt.PullRequest.Head.Repo.Name == w.GithubRepository {
			watch = i
		}
	}
	if watch == -1 {
		log.Printf("We are not interested in PRs (%s - %s)\n", *pr_evt.PullRequest.Head.User.Login, *pr_evt.PullRequest.Head.Repo.Name)
		return
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Config.Watches[0].GithubApiToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	if *pr_evt.PullRequest.State == "open" {
		status1 := &github.RepoStatus{State: &pending, TargetURL: &targetUrl, Description: &pendingDesc, Context: &appName}
		client.Repositories.CreateStatus(*pr_evt.PullRequest.Base.User.Login, *pr_evt.PullRequest.Base.Repo.Name, *pr_evt.PullRequest.Head.SHA, status1)


		s := RunJobByName(config.Config.Watches[watch].JenkinsJob, *pr_evt.Number, watch)

		log.Println("Completed job...")

		if s == true {
			log.Println("Returning Success")
			status2 := &github.RepoStatus{State: &success, TargetURL: &targetUrl, Description: &successDesc, Context: &appName}
			client.Repositories.CreateStatus(*pr_evt.PullRequest.Base.User.Login, *pr_evt.PullRequest.Base.Repo.Name, *pr_evt.PullRequest.Head.SHA, status2)
		} else {
			log.Println("Returning Error")
			status3 := &github.RepoStatus{State: &failure, TargetURL: &targetUrl, Description: &failureDesc, Context: &appName}
			client.Repositories.CreateStatus(*pr_evt.PullRequest.Base.User.Login, *pr_evt.PullRequest.Base.Repo.Name, *pr_evt.PullRequest.Head.SHA, status3)
		}
	} else {
			log.Printf("Ignoring PR event because state is %s.\n", *pr_evt.PullRequest.State)
	}
}

