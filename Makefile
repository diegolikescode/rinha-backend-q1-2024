lessgo:
	go run cmd/main.go

pg:
	docker-compose down
	docker-compose up --build

img:
	docker build -t rinha2024q1 .

all:
	docker build -t rinha2024q1 .
	docker-compose up --build

updown:
	docker-compose down --volumes --remove-orphans
	$(MAKE) all

comp:
	docker-compose up --build

t:
	./testes-local.sh
