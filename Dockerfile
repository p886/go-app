# Build stage
FROM golang:1.13 as build

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -v -o app

# Run stage
FROM alpine:latest
COPY --from=build go/src/app/ app/
CMD ["./app/app"]
