FROM golang:latest

WORKDIR /api
COPY ./ /api

RUN go build

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

EXPOSE 1323

CMD ["go", "run", "main.go"]
