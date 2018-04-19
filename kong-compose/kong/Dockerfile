FROM kong:0.13.0

RUN apk --no-cache add ca-certificates supervisor
RUN mkdir -p /var/log/supervisor
RUN mkdir /etc/consul.d
ADD supervisord.conf /etc/supervisord.conf
COPY --from=consul:1.0.6 /bin/consul      /usr/local/bin
COPY kong.json /etc/consul.d/

WORKDIR /etc/kong
COPY kong.conf .

CMD ["/usr/bin/supervisord"]