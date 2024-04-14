FROM golang:1.22 as build

WORKDIR /app

COPY . .

EXPOSE 8080

RUN make build-banner

ENTRYPOINT /app/server