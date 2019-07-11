package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	gitconfig "gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Config holds the runtime configuration
type Config struct {
	DestRepo string `required:"true"`
	Token    string `required:"true"`
}

func main() {

	cfg := Config{}
	err := envconfig.Process("plugin", &cfg)
	if err != nil {
		log.Fatalf("environment variable missing: %s", err)
	}

	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	_, err = r.CreateRemote(&gitconfig.RemoteConfig{
		Name: "dest",
		URLs: []string{cfg.DestRepo},
	})
	if err != nil {
		log.Fatalf("failed to add remote: %s", err)
	}

	err = r.Push(&git.PushOptions{
		RemoteName: "dest",
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: cfg.Token,
		},
	})
	if err != nil {
		log.Fatalf("failed to push: %s", err)
	}
	log.Info("Update done")

}
