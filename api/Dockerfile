FROM golang:latest

WORKDIR /api
COPY ./ /api

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 1323

ENTRYPOINT ["go", "run", "main.go"]
