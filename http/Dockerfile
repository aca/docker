FROM golang:1.15.2-alpine AS builder 

WORKDIR /src
COPY . .
RUN ["go", "build", "-o", "server"]

FROM alpine:latest
RUN apk add --no-cache ca-certificates

COPY --from=builder /src/server /bin/server
CMD ["server"]
