# Gitkins

Gitkins is a continous integration server that you can use from behind firewalls to interface between Github and Jenkins.

All help is appreciated and PRs welcome, please fork the repo and submit PRs with explanations as to what you want to do.

## Pre-Requisites
You need to:
 - Create a jenkins user and api that has access to run the jobs you want
 - Create a github api token from the user that will be able to recieve events and send statuses for the repos you want
 - See example config file below for other options that are needed.
 - You also need to create the hooks you want to receive events from (perhaps in a later version I will create/use programmatically created hooks)

## Docker Run
```
docker pull pafortin/gitkins:latest
docker run -d --name gitkins -p 8081:8081 pafortin/gitkins
```


## Example Gitkins-config.yaml file
```
port        : 8081
debug       : True
watches:
    - title              : github-events PR Checks
      githubAccount      : p4tin
      githubRepository   : github-events
      githubApiToken     : <your github api token here>
      jenkinsJob         : <job to run's name>
      jenkinsUrl         : http://localhost:32769/
      jenkinsUser        : admin
      jenkinsApiToken    : <your jenkins api token here>
```
