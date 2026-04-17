#FROM golang:1.22-alpine AS builder
#WORKDIR /build
#COPY . .
#RUN go mod tidy
#RUN go build -o server .
#
#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /build/server .
#COPY --from=builder /build/config ./config
#EXPOSE 8081
#CMD ["./server"]

FROM scratch

COPY . .

EXPOSE 8081

CMD ["/app"]