FROM golang:1.19 as build

WORKDIR /go/src/book

COPY . .

RUN mkdir -p bin &&\
    go mod download && go mod verify &&\
    GOOS=linux GOARCH=amd64 go build -o bin/

FROM golang:1.19

COPY --from=build /go/src/book/bin/book ./book

EXPOSE 8082
CMD ["./book"]