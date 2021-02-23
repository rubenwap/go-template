FROM golang:latest

WORKDIR /app/
COPY . .
RUN GOOS=linux go build -o example ./cmd/

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /opt/
COPY --from=0 /opt/example .
CMD ["./example"]