FROM golang:1.20-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /main cmd/rescheduler/main.go

FROM alpine:3
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]