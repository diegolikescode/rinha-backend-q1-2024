lessgo:
	go run cmd/main.go

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

pg:
	docker-compose -f compose-db.yml down --volumes --remove-orphans
	docker-compose -f compose-db.yml up --build

t:
	./testes-local.sh
