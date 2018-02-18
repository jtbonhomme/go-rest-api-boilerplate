# Go REST API Boilerplate

This repository is a skeleton for building production ready Golang REST API, with Swagger documentation.

# Installation

Assuming you have a working Go environment and GOPATH/bin is in your PATH.

## Dependencies

Install `gorilla/mux`:

```
$ go get github.com/gorilla/mux
```

## Live reload
[gin](https://github.com/codegangsta/gin) is a simple command line utility for live-reloading Go web applications. Just run `gin` in your app directory and your web app will be served with `gin` as a proxy. `gin` will automatically recompile your code when it detects a change. Your app will be restarted the next time it receives an HTTP request.

`gin` adheres to the "silence is golden" principle, so it will only complain if there was a compiler error or if you succesfully compile after an error.

`gin` is a breeze to install:

```
$ go get github.com/codegangsta/gin
```

Then verify that gin was installed correctly:

```
$ gin -h
```

Usage:
```
$ gin run -p 3000 -a 8000 main.go
```

Then verify that `gin` act correctly as a proxy to the api:

```
$ curl -X GET "http://localhost:3000/v1/people" -H "accept: application/json"
```

## Swagger

See the [go-swagger](https://github.com/go-swagger/go-swagger) page for full installation instructions.

`go-swagger` releases are distributed as binaries that are built from signed tags. It is published as github release, rpm, deb and docker image.

### Debian

```
$ echo "deb https://dl.bintray.com/go-swagger/goswagger-debian ubuntu main" | sudo tee -a /etc/apt/sources.list
```

### Homebrew

```
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```
