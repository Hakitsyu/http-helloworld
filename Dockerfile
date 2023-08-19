FROM golang:latest

WORKDIR /go/src/http-helloworld

COPY . .
RUN go mod download && go mod verify

RUN go install .

ENTRYPOINT /go/bin/http-helloworld

EXPOSE 8080
