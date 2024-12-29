FROM golang:1.23.4-alpine

WORKDIR /app

COPY go.mod ./
COPY . .

RUN go build -o main cmd/app/main.go

EXPOSE 8080

CMD ["./main"]
