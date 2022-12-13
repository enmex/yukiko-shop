FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/gateway/main.go

EXPOSE 8080
ENTRYPOINT ["./main"]
