FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .
RUN go get -d ./...
RUN CGO_ENABLED=0 go build -tags -a -installsuffix cgo -v -o main

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app/main .

EXPOSE 8080

CMD ["./main"]
