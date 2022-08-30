######## Start a builder stage #######
FROM golang:1.17-alpine as builder

WORKDIR /app
COPY ../.. .
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main main.go

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /opt/
COPY --from=builder /app/bin/main .
COPY --from=builder /app/build/ci/banner.txt .
USER root

CMD ["./main"]
