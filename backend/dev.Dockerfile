FROM golang:1.23

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/vektra/mockery/v2@v2.53.2

CMD ["air", "-c", ".air.toml"]
