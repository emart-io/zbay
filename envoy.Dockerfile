FROM envoyproxy/envoy-alpine:v1.14-latest

COPY ./envoy.yaml /etc/envoy/envoy.yaml
ADD ./cert/fullchain.cer ./cert/iyou.city.key /etc/
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

EXPOSE 443