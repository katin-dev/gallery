DB_DSN := "postgresql://gallery:gallery@localhost:5433/gallery?sslmode=disable"
IMAGE := gallery

build: 
	go build -o bin/gallery.exe main.go

build-img: 
	docker build -t ${IMAGE} .

run:
	bin/gallery.exe

up:
	docker-compose up -d

down:
	docker-compose down

migrate:
	goose --dir=migrations postgres ${DB_DSN} up
