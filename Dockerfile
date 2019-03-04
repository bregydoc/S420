FROM golang:1.11.4

WORKDIR /go/src/s420
COPY . .

RUN go get -v -d ./...
RUN go install -v ./...

CMD [ "launcher" ]
