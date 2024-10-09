EXECUTABLE=app

all: build run

run:
	./$(EXECUTABLE)

build:
	go build -o $(EXECUTABLE) cmd/main.go 

dockerCompose:
	docker-compose up

dockerStop:
	docker-compose down
	docker image rm online_song_library-app || true

swagInit:
	swag init -g cmd/main.go

clean:
	rm -rf ./$(EXECUTABLE)
#	rm -rf docs/
#	rm -rf logs/