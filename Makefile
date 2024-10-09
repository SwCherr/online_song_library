EXECUTABLE=app

all: build run

install:
	docker-compose up

uninstall:
	docker-compose down
	docker image rm online_song_library-app || true

run:
	./$(EXECUTABLE)

build:
	go build -o $(EXECUTABLE) cmd/main.go 

swagInit:
	swag init -g cmd/main.go

clean:
	rm -rf ./$(EXECUTABLE)
#	rm -rf docs/
#	rm -rf logs/