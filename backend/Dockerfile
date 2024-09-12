# Building the binary of the App
FROM golang:1.23.0 AS build

WORKDIR /backend

COPY go.mod go.sum ./

RUN go mod download

COPY *go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /slowers-backend

CMD ["/slowers-backend"]

