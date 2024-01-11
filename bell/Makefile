server:
	docker compose up --build

remove:
	docker stop bell-postgres-1
	docker stop bell-server-1
	docker rm bell-postgres-1
	docker rm bell-server-1

build:
	CGO_ENABLED=0 GOOS=linux go build -o bell

all:
	make build
	make server