FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o main ./cmd/image/main.go

EXPOSE 8082
ENTRYPOINT ["./main"]