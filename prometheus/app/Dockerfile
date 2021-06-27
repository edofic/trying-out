FROM golang:1.16-alpine as builder
COPY . /src
WORKDIR /src
RUN go build .

FROM alpine
WORKDIR /bin/
COPY --from=builder /src/app .
EXPOSE 2112
CMD /bin/app
