FROM golang:1.18

RUN apt-get update

ARG APP_NAME=blog-gin_golang_v177_mongo
RUN mkdir /$APP_NAME
COPY . /$APP_NAME
WORKDIR /$APP_NAME

RUN mv .env.staging .env
RUN go run ./db/migrate/migrate.go

CMD ["go","run","main.go"]
