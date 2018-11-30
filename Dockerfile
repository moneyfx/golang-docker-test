FROM golang

WORKDIR /go/src/hello-world

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["hello-world"]

