[![Maintainability](https://api.codeclimate.com/v1/badges/ad8dba8eeb423f1c7b3f/maintainability)](https://codeclimate.com/github/a4t/misudo/maintainability)

# misudo

Misudo is image uploader

## Description

## Installing

```
$ git clone git@github.com:a4t/misudo.git

or

$ wget -P /tmp https://github.com/a4t/misudo/releases/download/v0.0.1/misudo-v0.0.1-linux-amd64.zip
$ unzip misudo-v0.0.1-linux-amd64.zip
```

## Build

```
$ make deps
$ make

or

$ make cross-build
```

## Use

```
$ ./bin/misudo -p 8080
$ curl localhost:8080/ping
Pong
$ curl -XPOST -F output="response" -F w=300 -F file=@example/yagi.jpg localhost:8080/resize > example/yagi_resize.jpg
$ open example/yagi_resize.jpg
```
