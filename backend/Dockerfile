FROM golang:1.23

WORKDIR /app
COPY . .
RUN go build -o start-server

EXPOSE 5001

CMD ./start-server
