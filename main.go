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
	Force    bool   `optional`;
}

type DroneEnv struct {
	Branch	string `optional`;
}

func main() {

	cfg := Config{}
	err := envconfig.Process("plugin", &cfg)
	if err != nil {
		log.Fatalf("environment variable missing: %s", err)
	}

	drone := DroneEnv{}
	err = envconfig.Process("drone", &drone)
	if err != nil {
		log.Fatalf("failed to fetch drone env var: %s", err)
	}
	log.Info("Got drone env: %s", drone)

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

	ref := ""

	if cfg.Force == true {
		ref = "+"
		log.Info("force push requestet prepend: ", ref)
	}

	refspec := ref+drone.Branch+":"+drone.Branch
	pushconfig := gitconfig.RefSpec(refspec)
	log.Info("refspec: ", refspec)

	err = r.Push(&git.PushOptions{
		RefSpecs: []gitconfig.RefSpec{pushconfig},
		RemoteName: "dest",
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: cfg.Token,
		},
	})
	if err != nil {
		log.Fatalf("failed to push: %s", err)
	}
	log.Info("update done")

}
