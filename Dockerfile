FROM golang:1.19-alpine as build

WORKDIR /go/src/go-template

COPY . .

RUN make build

FROM scratch

COPY --from=build /go/src/go-template/bin/go-template ./app

EXPOSE 8080
CMD ["./App"]