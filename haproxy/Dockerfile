FROM haproxy:2.3.1
COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
COPY 404.http /etc/haproxy/errorfiles/404.http
COPY ssl/localhost/localhost.pem /etc/ssl/localhost/localhost.pem