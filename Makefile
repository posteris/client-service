swag:
	swag init --dir ./,./api/controllers --parseDependency --parseInternal

build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

start-test-env:
	test -f docker-compose.yml || wget -O  docker-compose.yml https://raw.githubusercontent.com/posteris/ci-database-starter/main/docker-compose.yml
	docker-compose up

bench:
	go test -bench 'Benchmark' ./...

test:
	go test ./...

