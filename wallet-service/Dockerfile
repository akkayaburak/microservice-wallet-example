FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]
