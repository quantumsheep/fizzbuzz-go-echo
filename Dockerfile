FROM golang:1.18 AS base

WORKDIR /app


FROM base AS dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]


FROM base AS prod

COPY . .

RUN go build -ldflags="-s -w" -o fizzbuzz-server .

CMD ["/app/fizzbuzz-server"]
