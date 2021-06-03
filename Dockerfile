FROM golang:1.16.4-alpine3.13 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -tags=nomsgpack -o main .

FROM alpine:3.13.5

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main /app/
CMD ["/app/main"]