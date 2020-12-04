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
