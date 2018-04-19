FROM postgres:9.5

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates wget && apt-get install -y  supervisor
RUN mkdir -p /var/log/supervisor
RUN mkdir /etc/consul.d
ADD supervisord.conf /etc/supervisord.conf
COPY --from=consul:1.0.6 /bin/consul      /usr/local/bin
COPY kong-database.json /etc/consul.d/

CMD ["/usr/bin/supervisord"]