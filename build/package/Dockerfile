FROM golang:latest as builder

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download

RUN go mod verify

COPY ../../cmd/ ./cmd/
COPY ../../internal/ ./internal
COPY ../../docs/ ./docs

RUN mkdir -p ./build/package

RUN CGO_ENABLED=0 GOOS=linux go build -o ./build/package/server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/build/package/server .

ENV ENVIRONMENT=production
# ENV GIN_MODE=release

EXPOSE 8080

CMD ["./server"]
