
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod ./

# RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base

COPY --from=builder /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
