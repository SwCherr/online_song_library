FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download


# createdb online_song_library
RUN [ "createdb", "online_song_library" ]

RUN [ "migrate", "-path", "./schema", "-database", "postgres://postgres:qwerty@postgres:5432/online_song_library?sslmode=disable", "up" ]

CMD [ "go", "run", "cmd/main.go" ]

# CMD [ "go", "build", "cmd/main.go", "-o", "app" ]



# FROM golang:1.23.1-alpine AS build

# WORKDIR /src

# COPY go.mod go.sum ./

# RUN go mod download

# COPY . .

# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

# RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# FROM alpine:3.20.3

# WORKDIR /app

# ENV GIN_MODE=release

# COPY --from=build /src/main /go/bin/migrate ./

# CMD ["/app/main"]
