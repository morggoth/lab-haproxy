FROM alpine AS ssl-builder
ENV DOMAIN localhost
RUN apk update && \
    apk add openssl && \
    mkdir /etc/ssl/${DOMAIN} && \
    openssl genrsa -out /etc/ssl/${DOMAIN}/${DOMAIN}.key 1024 && \
    openssl req -new -batch -key /etc/ssl/${DOMAIN}/${DOMAIN}.key \
        -out /etc/ssl/${DOMAIN}/${DOMAIN}.csr && \
    openssl x509 -req -days 365 -in /etc/ssl/${DOMAIN}/${DOMAIN}.csr \
        -signkey /etc/ssl/${DOMAIN}/${DOMAIN}.key \
        -out /etc/ssl/${DOMAIN}/${DOMAIN}.crt && \
    cat /etc/ssl/${DOMAIN}/${DOMAIN}.crt /etc/ssl/${DOMAIN}/${DOMAIN}.key \
        > /etc/ssl/${DOMAIN}/${DOMAIN}.pem

FROM haproxy:2.3.1
COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
COPY 404.http /etc/haproxy/errorfiles/404.http
COPY --from=ssl-builder /etc/ssl/localhost/localhost.pem /etc/ssl/localhost/localhost.pem