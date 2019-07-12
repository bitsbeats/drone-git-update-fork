package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
	gitconfig "gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

// Config holds the runtime configuration
type Config struct {
	DestRepo string `required:"true"`
	Token    string `required:"true"`
	Force    bool
}

// DroneEnv environment variables
type DroneEnv struct {
	Branch string
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
	log.Info("Got drone env:", drone)

	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("failed to open git repo: %s", err)
	}

	w, err := r.Worktree()
	if err != nil {
		log.Fatalf("failed to open worktree: %s", err)
	}

	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(drone.Branch),
	})

	_, err = r.CreateRemote(&gitconfig.RemoteConfig{
		Name: "dest",
		URLs: []string{cfg.DestRepo},
	})
	if err != nil {
		log.Fatalf("failed to add remote: %s", err)
	}

	err = r.Fetch(&git.FetchOptions{
		RemoteName: "dest",
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: cfg.Token,
		},
	})
	if err != nil {
		log.Fatalf("failed to fetch remote: %s", err)
	}

	refs, err := r.References()
	if err != nil {
		log.Fatalf("failed to get references: %s", err)
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		// The HEAD is omitted in a `git show-ref` so we ignore the symbolic
		// references, the HEAD
		if ref.Type() == plumbing.SymbolicReference {
			return nil
		}

		log.Info(ref)
		return nil
	})

	if cfg.Force == true {
		err = r.Push(&git.PushOptions{
			RefSpecs: []gitconfig.RefSpec{
			  "+refs/*:refs/*",
			  "+HEAD:refs/heads/HEAD",
		  },
		  RemoteName: "dest",
		  Auth: &http.BasicAuth{
			  Username: "abc123", // yes, this can be anything except an empty string
			  Password: cfg.Token,
		  },
	  })
	  if err != nil {
		  log.Warnf("failed to push: %s", err)
	  }
	} else {
		err = r.Push(&git.PushOptions{
			RefSpecs: []gitconfig.RefSpec{
			  "refs/*:refs/*",
			  "HEAD:refs/heads/HEAD",
		  },
		  RemoteName: "dest",
		  Auth: &http.BasicAuth{
			  Username: "abc123", // yes, this can be anything except an empty string
			  Password: cfg.Token,
		  },
	  })
	  if err != nil {
		  log.Warnf("failed to push: %s", err)
	  }
	}

	log.Info("update done")

}
