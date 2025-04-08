FROM node:18 AS frontend-build

WORKDIR /app

COPY /frontend .

RUN npm clean-install
RUN npm run build


FROM golang:1.23 AS backend-build

WORKDIR /app

COPY /backend/go.mod /backend/go.sum ./
RUN go mod download

COPY /backend .

RUN go build -o start-server


FROM debian:12.9-slim

RUN useradd -ms /bin/bash server

WORKDIR /app

COPY --chown=server:server --from=frontend-build /app/dist ./client/dist
COPY --chown=server:server --from=backend-build /app/start-server .
COPY --chown=server:server /backend/database/psql/schema.sql /backend/database/psql/functions.sql ./database/psql/

USER server

ENTRYPOINT [ "./start-server" ]
