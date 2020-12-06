FROM golang:alpine as go-builder
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main .

FROM alpine
EXPOSE 8000
COPY --from=go-builder /app/main /app/main
CMD ["/app/main"]