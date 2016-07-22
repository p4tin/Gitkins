package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	GithubToken string
	JenkinsUsername string
	JenkinsApiToken string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	GithubToken = os.Getenv("GITHUBTOKEN")
	JenkinsUsername = os.Getenv("JENKINSUSERNAME")
	JenkinsApiToken = os.Getenv("JENKINSAPITOKEN")
}
