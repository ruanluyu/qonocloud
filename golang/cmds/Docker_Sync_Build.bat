docker build --tag qonocloud:sync -f %~dp0..\Dockerfile.sync %~dp0..\
docker image prune -f
