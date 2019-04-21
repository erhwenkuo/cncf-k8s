From busybox:glibc

RUN mkdir -p /root/workplace
COPY ./app2 /root/workplace/app2
COPY ./server.crt /root/workplace/server.crt
COPY ./server.key /root/workplace/server.key
RUN chmod +x /root/workplace/app2

WORKDIR /root/workplace
ENTRYPOINT ["/root/workplace/app2"]

