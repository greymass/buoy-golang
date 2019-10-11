buoy
====

Dumb POST -> WebSocket forwarder

Run with docker
---------------

```
$ docker build .
...
<container id>
$ docker run -e ADDR='0.0.0.0:8080' -p 8080:8080 --name buoy <container id>
```

Run with docker-compose
---------------

```
$ docker-compose build
$ docker-compose up -d
```
