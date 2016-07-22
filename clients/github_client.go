package clients

import (
	"golang.org/x/oauth2"
	"github.com/google/go-github/github"

	"github.com/p4tin/Gitkins/config"
)

var client *github.Client

func init() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.GithubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client = github.NewClient(tc)
}


