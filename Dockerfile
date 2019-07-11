FROM busybox

ADD ./drone-git-update-fork /drone-git-update-fork 

ENTRYPOINT ["/drone-git-update-fork"]