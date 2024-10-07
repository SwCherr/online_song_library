all: clean build mocks test gcov

# it`s temporary goal
start: deleteDB createDB run

run: clean swagInit
	go run cmd/main.go

# createFileMigration:
# 	migrate create -ext sql -dir schema -seq init

createDB:
	createdb online_song_library
	migrate -path ./schema -database "postgres://uliakungurova:qwerty@localhost:5432/online_song_library?sslmode=disable" up

deleteDB:
	migrate -path ./schema -database "postgres://uliakungurova:qwerty@localhost:5432/online_song_library?sslmode=disable" down || true
	dropdb -f --if-exists -e online_song_library

swagInit:
	swag init -g cmd/main.go

clean:
	rm -rf docs/


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