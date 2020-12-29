# Rate Limit

[![Lint](https://github.com/leozhantw/rate-limit/workflows/Lint/badge.svg)](https://github.com/leozhantw/rate-limit/actions?query=workflow%3ALint)
[![Test](https://github.com/leozhantw/rate-limit/workflows/Test/badge.svg)](https://github.com/leozhantw/rate-limit/actions?query=workflow%3ATest)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/leozhantw/rate-limit/blob/master/LICENSE)

Rate limit implementation by Go

## To Start Developing
If you want to build this service right away there are two options:

**You have a working [Go environment](https://golang.org/doc/install).**
```shell
git clone https://github.com/leozhantw/rate-limit.git
cd rate-limit
go run cmd/server/server.go
```

**You have a working [Docker environment](https://docs.docker.com/engine/).**
```shell
git clone https://github.com/leozhantw/rate-limit.git
cd rate-limit
docker build -t rate-limit .
docker run -p 3000:3000 rate-limit:latest
```
