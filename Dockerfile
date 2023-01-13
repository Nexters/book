FROM golang:1.19 as build

WORKDIR /go/src/go-template

COPY . .

RUN mkdir -p bin &&\
    go mod download && go mod verify &&\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/

FROM scratch
COPY --from=build /go/src/go-template/bin/go-template ./app

EXPOSE 8080
CMD ["./app"]