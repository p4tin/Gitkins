package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Version = "0.0.1"
)

type Job struct {
	GithubAccount    string `yaml:"githubAccount"`
	GithubApiToken   string `yaml:"githubApiToken"`
	GithubRepository string `yaml:"githubRepository"`
	JenkinsJob       string `yaml:"jenkinsJob"`
	JenkinsUrl       string `yaml:"jenkinsUrl"`
	JenkinsUser      string `yaml:"jenkinsUser"`
	JenkinsApiToken  string `yaml:"jenkinsApiToken"`
}

type configStruct struct {
	Port    string `yaml:"port"`
	Debug   bool   `yaml:"debug"`
	Watches []Job  `yaml:"watches"`
}

var Config configStruct

func init() {
	configFile := flag.String("config", "./Gitkins-config.yaml", "Config file")
	flag.Parse()

	dat, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Println("Could not load yaml config file.")
	}

	err = yaml.Unmarshal(dat, &Config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
