FROM golang:1.9-alpine

ENV TZ=Asia/Shanghai

RUN  apk add --no-cache mongodb curl bash ca-certificates && \

 rm /usr/bin/mongoperf

RUN go build -o berk && chmod +x berk

ENV TZ=Asia/Shanghai
ADD ./berk /
ADD ./certs/* /certs/
ADD ./ssh/*  ~/.ssh/
ADD ./nginx.conf /etc/nginx/conf.d
EXPOSE  80 443 27017 28017
VOLUME /data/db

COPY ./scripts/entrypoint.sh /root
ENTRYPOINT [ "/root/entrypoint.sh","./berk"]
CMD [ "mongod" ]
