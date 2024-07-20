FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/stresstest

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/stresstest .

ENTRYPOINT [ "./stresstest" ]