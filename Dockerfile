FROM alpine:latest

RUN echo "Asia/shanghai" >> /etc/timezone

COPY ./main /bin/kk-uuid

RUN chmod +x /bin/kk-uuid

ENV KK_NAME kk.uuid.
ENV KK_ADDR 127.0.0.1:87

CMD kk-uuid $KK_NAME $KK_ADDR
