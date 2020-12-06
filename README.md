# lab-haproxy

Playground for HAProxy experiments.

```bash
# Run to start this lab
docker-compose up -d

# Run to send some requests to the Lab
ab -n 1000 http://localhost:8080/
```

Then go to <http://localhost:8080> to get access to the Lab's endpoint.

## Statistics

<http://localhost:8080/custom-uri> - page with HAProxy statistics.

Credentials:

- user: admin
- pass: admin

## SSL

How to create self-signed certificate:

```bash
DOMAINE="domain.name"
mkdir /etc/ssl/${DOMAIN}
openssl genrsa -out /etc/ssl/${DOMAIN}/${DOMAIN}.key 1024
openssl req -new -key /etc/ssl/${DOMAIN}/${DOMAIN}.key \
    -out /etc/ssl/${DOMAIN}/${DOMAIN}.csr
openssl x509 -req -days 365 -in /etc/ssl/${DOMAIN}/${DOMAIN}.csr \
    -signkey /etc/ssl/${DOMAIN}/${DOMAIN}.key \
    -out /etc/ssl/${DOMAIN}/${DOMAIN}.crt

cat /etc/ssl/${DOMAIN}/${DOMAIN}.crt /etc/ssl/${DOMAIN}/${DOMAIN}.key \
    | tee /etc/ssl/${DOMAIN}/${DOMAIN}.pem
```

<https://localhost:8443/termination> - App instances configured with SSL termination on HAProxy.
<https://localhost:8444> - App instances configured with SSL pass-through.

NB: don't forget flag _-k_ for _curl_ to avoid error with self-signed certificate
