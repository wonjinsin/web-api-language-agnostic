# wonjinsin/go-web-boilerplate(pikachu)

Simple rest api server with [Echo framework](https://github.com/labstack/echo)

[![License MIT](https://img.shields.io/badge/License-MIT-green.svg)](http://opensource.org/licenses/MIT)
[![DB Version](https://img.shields.io/badge/DB-Mysql-blue)](https://www.mysql.com/)
[![Go Report Card](https://goreportcard.com/badge/github.com/StarpTech/go-web)](https://goreportcard.com/report/github.com/wonjinsin/go-web-boilerplate)

## Features

- [Gorm](https://github.com/go-gorm/gorm) : For mysql ORM
- [Zap](https://github.com/uber-go/zap) : Go leveled Logging
- [Viper](https://github.com/spf13/viper) : Config Setting
- [Makefile]() : go build, test, vendor using Make

## Project structure

Clean Architecture with DDD(Domain Driven Design) pattern

## Getting started

### Set infra

```
$ docker-compose -f infra-docker.yml up -d
```

### Initial action

```
$ make all && make build && make start
```

### Migration initial data

```
$ make migrate up ENV=local
```

If you have error when Init Please use below command and do Inital action again

```
make clean
```

## MakeFile Command

### migrate up

```
# e.g make migrate up ENV=local
$ make migrate up ENV=${ENV}
```

### migrate down

```
# e.g make migrate down ENV=local
$ make migrate down ENV=${ENV}
```

### Build vendors

```
$ make vendor
```

### Build and start

```
$ make build && bin/pikachu
```

### Test

```
$ make vet && make fmt && make lint && make test
// or
$ make test-all
```

### Clean

```
$ make clean
```

## Docs

### Specification

https://github.com/upsidr/coding-test/blob/main/web-api-language-agnostic/README.ja.md

### Api docs

- Use this with post man, swagger etc

  [open api](openapi/init.yml)

## How to use

1. Set initial datas at [Getting Started](#getting-started)

2. Get access token with /auths/signin api, mock user is already created

3. Put access token to Bearer token, and call other apis
