FROM alpine:latest
MAINTAINER jiangpengfei <jiangpengfei12@gmail.com>

COPY ./build/cola /root/cola

CMD ["/root/cola"]
