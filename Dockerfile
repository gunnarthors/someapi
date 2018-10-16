FROM golang:1.11 AS builder

WORKDIR $GOPATH/src/github.com/someapi
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /main .

FROM scratch
COPY --from=builder /main ./
EXPOSE 8080
ENV ENVVAR="TEST"
ENTRYPOINT ["./main"]
