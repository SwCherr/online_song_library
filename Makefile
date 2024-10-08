EXECUTABLE=app

all: run

run: build
	./$(EXECUTABLE)

build: swagInit
	go build -o $(EXECUTABLE) cmd/main.go 

# test:
# 	go test --coverprofile=c.out ./pkg/handler/

# gcov:
# 	go tool cover -html=c.out

# # dependencies
# mocks: clean
# 	mockgen -source=pkg/service/service.go -destination=mocks/service/mock_service.go

dockerCompose:
	docker-compose up

dockerStop:
	docker-compose down

swagInit:
	swag init -g cmd/main.go

clean:
	rm -rf ./$(EXECUTABLE)
	rm -rf docs/
	rm -rf logs/
#	rm -rf mocks/


# LOCAL DATABASE
# createFileMigration:
# 	migrate create -ext sql -dir schema -seq init

# createDB:
# 	createdb online_song_library || true
# 	migrate -path ./schema -database "postgres://uliakungurova:qwerty@localhost:5432/online_song_library?sslmode=disable" up

# deleteDB:
# 	migrate -path ./schema -database "postgres://uliakungurova:qwerty@localhost:5432/online_song_library?sslmode=disable" down || true
# 	dropdb -f --if-exists -e online_song_library

	
# .PHONY: clean