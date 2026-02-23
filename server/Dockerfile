FROM golang:1.25

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /banana-split-server

EXPOSE 8001

CMD ["/banana-split-server"]