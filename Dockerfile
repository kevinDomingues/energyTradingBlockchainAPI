FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/main /main

CMD ["/main"]

EXPOSE 5000
