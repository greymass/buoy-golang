NOT PRODUCTION READY
---

The golang version of Buoy has lagged behind in functionality compared to that of the nodejs version. It is recommended for the time being that the nodejs version is used until development can continue on this version.

https://github.com/greymass/buoy-nodejs

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
