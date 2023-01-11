FROM registry-gitlab.nexon.com/infraleadingtech/images/go:1.19-alpine

WORKDIR /go/src/go-template

COPY . .

RUN make build 

EXPOSE 8080
CMD ["./bin/app"]