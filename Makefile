all: clean build mocks test gcov

# DOCKER
dockerRun:
#	docker run -it online_song_library
	docker run online_song_library-app

# dockerBuild:
# 	docker build . -t online_song_library:latest

dockerCompose:
	docker-compose up

dockerStop:
	docker-compose down

# APP
start: createDB run

build:
	go build cmd/main.go -o app

run:
	./app

# run: 
# 	go run cmd/main.go

runSwag: clean swagInit run

# createFileMigration:
# 	migrate create -ext sql -dir schema -seq init

createDB:
#	createdb online_song_library
	createdb -p 5432 -h postgres -E LATIN1 -e online_song_library
	migrate -path ./schema -database "postgres://postgres:qwerty@postgres:5432/online_song_library?sslmode=disable" up

# Создать базу demo на сервере eden, порт 5000, с кодировкой LATIN1 можно так:
# createdb -p 5432 -h postgres -E LATIN1 -e online_song_library
# createdb -p 5000 -h eden -E LATIN1 -e demo

# CREATE DATABASE demo ENCODING 'LATIN1';


deleteDB:
	migrate -path ./schema -database "postgres://postgres:qwerty@postgres:5432/online_song_library?sslmode=disable" down || true
	dropdb -f --if-exists -e online_song_library

# SWAGGER
swagInit:
	swag init -g cmd/main.go

clean:
	rm -rf docs/
	rm -rf ./app


# build:
# 	go build cmd/main.go

# run:
# 	./main

# test:
# 	go test --coverprofile=c.out ./pkg/handler/

# gcov:
# 	go tool cover -html=c.out

# clean: 
# 	rm -rf c.out
# 	rm -rf main
# #	rm -rf mocks/

# # dependencies
# mocks: clean
# 	mockgen -source=pkg/service/service.go -destination=mocks/service/mock_service.go

# postgres: dockerPosgres generateDB
# #	docker exec -it ed9f79168839 /bin/bash
	
# dockerPosgres:
# 	docker pull postgres
# 	docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

# # createFileMigration:
# # 	migrate create -ext sql -dir ./schema -seq init
	
# generateDB:
# 	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

# cleanDB:
# 	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down

# .PHONY: clean