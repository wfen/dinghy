# build a main binary using the golang container
FROM golang:1.11 as builder
MAINTAINER Ward Fenton <ward.fenton@gmail.com>

WORKDIR /home/ward/code/dory/experiments/tcp-echo-server/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# copy the main binary to a separate container based on ubuntu
FROM ubuntu:18.04
WORKDIR /bin/
COPY --from=builder /home/ward/code/dory/experiments/tcp-echo-server/main .
ENTRYPOINT [ "/bin/main" ]
CMD [ "9000", "hello" ]
EXPOSE 9000
