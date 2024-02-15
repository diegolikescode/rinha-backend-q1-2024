FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main cmd/main.go

EXPOSE 6969

CMD ["./main"]
