FROM golang:1.20.11-alpine as builder

RUN apk update && apk add git make

WORKDIR /app

COPY . .

COPY go.* ./

RUN go mod download

RUN go mod verify

RUN CGO_ENABLED=0 go build -o MyApp ./cmd/

RUN chmod +x /app/MyApp

EXPOSE 80

CMD ["/app/MyApp"]
