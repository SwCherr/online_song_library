FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download


# createdb online_song_library
RUN [ "createdb", "online_song_library" ]

RUN [ "migrate", "-path", "./schema", "-database", "postgres://postgres:qwerty@postgres:5432/online_song_library?sslmode=disable", "up" ]

CMD [ "go", "run", "cmd/main.go" ]

# CMD [ "go", "build", "cmd/main.go", "-o", "app" ]