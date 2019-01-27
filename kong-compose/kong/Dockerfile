FROM kong:0.14.1

RUN apk --update --no-cache add ca-certificates supervisor dnsmasq curl
RUN echo 'conf-dir=/etc/dnsmasq.d/,*.conf' > /etc/dnsmasq.conf
RUN echo "user=root" >> /etc/dnsmasq.conf
RUN echo "server=/consul/127.0.0.1#8600" > /etc/dnsmasq.d/consul.conf
RUN mkdir -p /var/log/supervisor
RUN mkdir /etc/consul.d
ADD supervisord.conf /etc/supervisord.conf
COPY --from=consul:1.0.6 /bin/consul      /usr/local/bin
COPY --from=hashicorp/consul-template:alpine /bin/consul-template   /usr/local/bin
COPY kong.json /etc/consul.d/

COPY kong_conf.ctmpl /tmp
# WORKDIR /etc/kong
# COPY kong.conf .

CMD ["/usr/bin/supervisord"]