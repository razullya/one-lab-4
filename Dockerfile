FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY ./ /go/src/app/
RUN go build -o main main.go
FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/
CMD [ "./main" ]