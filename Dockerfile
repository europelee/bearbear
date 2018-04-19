FROM golang:1.10-alpine as builder
RUN	apk	--no-cache	add	git	ca-certificates
WORKDIR	/opt/BearBear
COPY  main.go .
RUN	go	get	-d	-v	github.com/gin-gonic/gin \
&&	CGO_ENABLED=0	GOOS=linux	go	build	-a	-installsuffix	cgo	-o	bearbear	.

FROM	alpine:3.4	as	out
RUN apk --no-cache add ca-certificates bind-tools supervisor
RUN mkdir -p /var/log/supervisor
RUN mkdir /etc/consul.d
ADD supervisord.conf /etc/supervisord.conf
COPY --from=0 /opt/BearBear/bearbear	/usr/local/bin
COPY --from=consul:1.0.6 /bin/consul      /usr/local/bin
COPY bearbear.json /etc/consul.d/
CMD ["/usr/bin/supervisord"]