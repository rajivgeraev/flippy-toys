FROM golang:1.23.4-alpine

WORKDIR /app
COPY . .
RUN go build -o api cmd/app/main.go

EXPOSE 8080
CMD ["./api"]