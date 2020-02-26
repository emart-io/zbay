FROM envoyproxy/envoy-alpine:latest

COPY ./envoy.yaml /etc/envoy/envoy.yaml
ADD ./fullchain.pem ./privkey.pem /etc/
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

EXPOSE 443