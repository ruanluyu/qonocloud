docker build --tag qonocloud:latest -f %~dp0..\Dockerfile.final %~dp0..\
docker image prune -f
