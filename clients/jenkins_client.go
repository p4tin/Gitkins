package clients

import (
	"log"
	"github.com/bndr/gojenkins"
	"strconv"

	"github.com/p4tin/Gitkins/config"
	"time"
)

func RunJobByName(name string, id int, watch int) bool {
	log.Printf("Running Job %s\n", name)

	jenkins, err := gojenkins.CreateJenkins(config.Config.Watches[watch].JenkinsUrl, config.Config.Watches[watch].JenkinsUser, config.Config.Watches[watch].JenkinsApiToken).Init()

	if err != nil {
		log.Printf("Could not connect to Jenkins at:", config.Config.Watches[watch].JenkinsUrl)
		return false
	}

	job, err := jenkins.GetJob(name)
	if err != nil {
		log.Printf("Job %s not found\n", name)
		return false
	}
	params := map[string]string {
		"PR": strconv.Itoa(id),
		"URL":   "Test1",
	}
	stat, err := job.InvokeSimple(params)
	if err != nil {
		log.Println("Build returned error: " + err.Error())
		return false
	}
	log.Printf("%b last job invokde response", stat)
	IsRunning := true
	for IsRunning {
		IsRunning, err = job.IsQueued()
		if err != nil {
			log.Println("Could not find out if the build was still running or not.")
			return false
		}
		if !IsRunning {
			IsRunning, err = job.IsRunning()
			if err != nil {
				log.Println("Could not find out if the build was still running or not.")
				return false
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	b, err := job.GetLastBuild();
	if err != nil {
		log.Println("Could not get the last build's status")
		return false
	}
	log.Printf("Build ID %s complete with status %s\n", b.Raw.ID, b.Raw.Result)

	return true
}
