FROM golang:latest

USER api

WORKDIR /api
COPY ./ /api

RUN go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 1323

ENTRYPOINT ["go", "run", "main.go"]
