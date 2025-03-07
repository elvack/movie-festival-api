FROM golang:1.24

RUN apt-get update

ARG APP_NAME=movie-festival-api
RUN mkdir /$APP_NAME
COPY . /$APP_NAME
WORKDIR /$APP_NAME

RUN go run database/migrate/migrate.go -t up
RUN go mod download
RUN go build -o movie_festival_api .

CMD ["./movie_festival_api"]
