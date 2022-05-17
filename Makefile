build: go-clean go-build

up: docker-dev-up

run: build go-run

test: go-test
#docker-test-up

clean: go-clean

go-clean:
	go clean
	rm -rf main

go-build:
	go build ./src/main.go

go-run:
	go run ./src/main.go

go-test:
	go test -v ./src/...

docker-dev-up:
	docker-compose --file ./docker/development/docker-compose.yml up

docker-test-up:
	docker-compose --file ./docker/testing/docker-compose.yml build --force
	docker-compose --file ./docker/testing/docker-compose.yml up --abort-on-container-exit
