# Refer to https://hub.docker.com/_/alpine/tags
FROM alpine:3.18.3

MAINTAINER Vuong Le lelordoftech@gmail.com

# easy-rsa
RUN apk update \
 && apk add --update-cache \
    easy-rsa \
 && rm -rf /var/cache/apk/*

WORKDIR /opt

# Needed by scripts
ENV OPENVPN=/etc/openvpn
ENV CA_SERVER=/etc/ca_server

ADD assets/generate_ca_and_server_certs.sh /opt/scripts/generate_ca_and_server_certs.sh
ADD assets/vars.template /opt/scripts/vars.template
ADD assets/start.sh /opt/start.sh
RUN chmod +x /opt/scripts/generate_ca_and_server_certs.sh /opt/start.sh

ADD openvpn-web-ui.tar.gz /opt/openvpn-gui/
RUN rm -f /opt/openvpn-gui/data.db
ADD assets/app.conf /opt/openvpn-gui/conf/app.conf

EXPOSE 8080/tcp

CMD /opt/start.sh
