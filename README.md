# GolangApplicationDevelopment
Golang Application Development - KBTU course repository.

made by Tursynbay Dinmukhamed

to start server run following command in greenlight/
```shell
go run ./cmd/api 
```

connections get env variable .env
```shell
source .env 
echo $GREENLIGHT_DB_DSN
```

connections to DB
```shell
docker exec -it $DOCKER_DB bash
psql -U greenlight -d greenlight
```

create migration to DB
```shell
migrate create -seq -ext=.sql -dir=./migrations name
```

migrate up/down(apply/revert)
```shell
migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```
