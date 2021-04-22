FROM envoyproxy/envoy-alpine:v1.18-latest
# RESTful接口专用
COPY ./envoy.yaml /etc/envoy/envoy.yaml
ADD ./service/descriptor.pb /etc/
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

EXPOSE 8080