FROM golang:1.25

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.* ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/banana ./server/cmd/

EXPOSE 8001

CMD ["/app/bin/banana"]