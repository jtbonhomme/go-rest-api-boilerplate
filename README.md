# Go REST API Boilerplate

This repository is a skeleton for building production ready Golang REST API, with Swagger documentation.

We will use gorilla/mux and go-swagger packages.

![](images/gorilla.png) ![](images/swagger.png)


# Installation

Assuming you have a working Go environment and GOPATH/bin is in your PATH.

## Dependencies

Install `gorilla/mux`:

```
$ go get github.com/gorilla/mux
```

Install CORS middleware `rs/cors`:
```
$ go get github.com/rs/cors
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

Start gin proxy to listen on port 3000 and send request to proxied app listening on port 8000:
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

### Generate swagger

```
$ swagger generate spec -o swagger-api/v1/swagger.json && swagger validate swagger-api/v1/swagger.json
```

### Serve swagger documentation

#### Swagger Flavor

```
$ swagger serve --flavor=swagger --port=6060 swagger-api/v1/swagger.json
```

This command will open your browser on the URL `petstore.swagger.io` (the official [swagger](https://swagger.io/swagger-ui/) live-demo exploration tool). This server allow you to explore any swagger json file.

![](images/swagger-flavor.png)

You can host your own swagger-ui server too: https://swagger.io/docs/swagger-tools/#download-33

#### Redoc flavor

```
$ swagger serve --flavor=redoc --port=6060 swagger-api/v1/swagger.json
```
This command will open your browser on the URL `localhost:6060/docs`, no need of an internet connexion to use it.

![](images/redoc-flavor.png)

# REST API

## Content of the package

```
├── main.go
├── router
│   ├── logger.go
│   ├── routes.go
│   └── router.go
├── db
│   └── people.go
├── handlers
│   ├── people.go
│   └── response.go
├── model
│   └── people.go
└── swagger-api
    └── v1
        └── swagger.json
```

### main.go

### router

### handlers

### db

### models

### swagger-api/v1

# References

The following references helped me :

* https://github.com/gorilla/mux
* https://github.com/rs/cors
* http://www.gorillatoolkit.org/pkg/mux
* https://goswagger.io/
* https://github.com/kkamdooong/go-restful-api-example
* https://thenewstack.io/make-a-restful-json-api-go/
* https://github.com/corylanou/tns-restful-json-api

# Todo

- [ ] Error handling for any system call
- [X] Add correct Content-type headers
- [ ] Handle correctly Accept headers
- [ ] Handle concurrency
- [ ] Add pagination for collections/models getters
- [ ] Http StatusCode handling with consistent http codes implementation
- [ ] Versionning with version.go or local json file (pathPrefix)
- [ ] Full swagger documentation
- [ ] Aggregate multiple miroservices API in a single Swagger end point
- [ ] Authentication
- [ ] Handle [eTags](http://en.wikipedia.org/wiki/HTTP_ETag)
- [ ] Add environment variable such as listening port
- [ ] Production ready logs
- [ ] Unitary and system Tests
- [ ] Rate limit
- [ ] gRPC example/version