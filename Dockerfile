FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o my-book-app ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/my-book-app .

EXPOSE 8080

CMD ["./my-book-app"]
