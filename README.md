```shell
export $(grep -v '^#' .env | xargs -0)
```

```shell
docker-compose up -d --no-deps --build
```