From busybox:glibc

RUN mkdir -p /root/workplace
COPY ./app3 /root/workplace/app3
RUN chmod +x /root/workplace/app3

WORKDIR /root/workplace
ENTRYPOINT ["/root/workplace/app3"]

