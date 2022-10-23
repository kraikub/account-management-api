FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /account-management-api/src

COPY . .

RUN go mod tidy


RUN go build -o app ./api/v1/cmd/main.go

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /account-management-api/src/app .

CMD ["./app"]