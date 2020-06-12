#FROM golang:1.14.3-alpine
#WORKDIR /go/emart
#ADD ./ ./
#ENV GOPROXY=https://goproxy.cn
#RUN go build -o zbay internal/cmd/*.go

FROM alpine:3.5
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#RUN apk --no-cache add ca-certificates
COPY bundles/zbay /usr/local/bin
#COPY --from=0 /go/emart/zbay /usr/local/bin
ENTRYPOINT ["/usr/local/bin/zbay"]

EXPOSE 50051 9090