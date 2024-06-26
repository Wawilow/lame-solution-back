Here is simple website backend, set up your telegram token to the .env file to send contact messages

Local run
```shell
export $(grep -v '^#' .env | xargs -0)
go run cmd/server/main.go
```

Docker run (docker-compose file and everything for k8s deploying is in the head repo)
```shell
docker-compose up -d --no-deps --build
```
