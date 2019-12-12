FROM envoyproxy/envoy-alpine:latest

#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD ./bundles/zbay /usr/local/bin
ENTRYPOINT ["/usr/local/bin/zbay"]

EXPOSE 50051 9090