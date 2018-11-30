FROM golang

WORKDIR /go/src/hello-world

COPY . .

#RUN go run hello-world.go

RUN go build hello-world.go

CMD ["./hello-world"]

