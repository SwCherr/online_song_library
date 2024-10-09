EXECUTABLE=app

all: build run

run:
	./$(EXECUTABLE)

build: swagInit
	go build -o $(EXECUTABLE) cmd/main.go 

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