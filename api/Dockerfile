FROM golang:latest as builder

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

COPY ./ /willers-api
WORKDIR /willers-api
RUN go build -o api 

# --

FROM golang:latest 
COPY --from=builder /willers-api/api /bin

WORKDIR /bin
EXPOSE 1323

ENTRYPOINT ["/bin/api"]

