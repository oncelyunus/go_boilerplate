FROM golang:1.19-buster
WORKDIR /app


COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o main .
ENTRYPOINT ["/app/main"]