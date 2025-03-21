FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go install github.com/vektra/mockery/v2@v2.53.2

CMD mockery ; air -c .air.toml
