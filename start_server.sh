#!/bin/sh

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate -path ./schema -database 'postgres://postgres:qwerty@postgres:5432/online_song_library?sslmode=disable' up

go run cmd/main.go