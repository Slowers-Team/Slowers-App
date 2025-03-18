FROM node:18 AS frontend-build

WORKDIR /app

COPY /frontend .

RUN npm clean-install
RUN npm run test
RUN npm run build


FROM golang:1.23 AS backend-build

WORKDIR /app

COPY /backend/go.mod /backend/go.sum ./
RUN go mod download

COPY /backend .
RUN go install github.com/vektra/mockery/v2@v2.53.2
RUN mockery
RUN go test ./...

RUN go build -o start-server


FROM debian:12.9-slim

RUN useradd -ms /bin/bash server

WORKDIR /app

COPY --chown=server:server --from=frontend-build /app/dist ./client/dist
COPY --chown=server:server --from=backend-build /app/start-server .

USER server
EXPOSE 5001
ENTRYPOINT [ "./start-server" ]
