FROM golang:1.21.7

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "cmd/main.go", "-config=./config.toml"]