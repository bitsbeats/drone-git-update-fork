# drone-git-update-fork

[![Build Status](https://cloud.drone.io/api/badges/bitsbeats/drone-git-update-fork/status.svg)](https://cloud.drone.io/bitsbeats/drone-git-update-fork)
[![Gitter chat](https://badges.gitter.im/drone/drone.png)](https://gitter.im/drone/drone)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Report](https://goreportcard.com/badge/github.com/bitsbeats/drone-git-update-fork)](https://goreportcard.com/badge/github.com/bitsbeats/drone-git-update-fork)

Drone plugin to update a remote git repository, e.g. a fork.

## Build
```console
go build .
```

## Run

```console
PLUGIN_DESTREPO=https://github.com/foobar/destination-repo.git PLUGIN_TOKEN=<github token> DRONE_BRANCH=master drone-git-update-fork
```

## Example drone step

```console
---
name: default
kind: pipeline

steps:

  - name: update git fork
    image: bitsbeats/drone-git-update-fork
    settings:
      destrepo: https://github.com/foobar/destination-repo.git
      token:
        from_secret: github_token
      force: true

```
