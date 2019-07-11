FROM alpine

RUN apk update && apk add ca-certificates
ADD ./drone-git-update-fork /drone-git-update-fork 

ENTRYPOINT ["/drone-git-update-fork"]
