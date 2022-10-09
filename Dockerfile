FROM golang:1.18-alpine as builder
WORKDIR /usr/app
# copy go.mod and go.sum 
ADD go.* ./
RUN go mod download
# copy src
ADD *.go ./
COPY providers/ providers/
COPY servers/  servers/
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bot .

FROM alpine:3
COPY --from=builder /usr/app/bot .
# executable
ENTRYPOINT [ "./bot" ]
# arguments that can be overridden
CMD [ "amqp" ]