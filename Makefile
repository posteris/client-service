swag:
	swag init --dir ./,./api/controllers --parseDependency --parseInternal

build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run
