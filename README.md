# chi-param
URL parameters helper for [Сhi Router](https://github.com/go-chi/chi).

[![Go Report Card](https://goreportcard.com/badge/github.com/oceanicdev/chi-param?style=flat-square)](https://goreportcard.com/report/github.com/oceanicdev/chi-param)
[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/oceanicdev/chi-param/master/LICENSE)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/oceanicdev/chi-param)
[![Release](https://img.shields.io/github/release/oceanicdev/chi-param.svg?style=flat-square)](https://github.com/oceanicdev/chi-param/releases/latest)

## Install

```bash
go get -u github.com/oceanicdev/chi-param
```

And then import the package in your code:

```go
import "github.com/oceanicdev/chi-param"
```

### Example

An example described below is one of the use cases.

```go
package main

import (
	"github.com/go-chi/chi"
	"github.com/oceanicdev/chi-param"
	"log"
	"net/http"
)

func main() {
	r := chi.NewMux()
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		// call http://localhost:8080/1?code=1&code2
		id, _ := param.Int(r, "id")                // returns value from path
		code, _ := param.QueryInt(r, "code")       // returns first value
		codes, _ := param.QueryIntArray(r, "code") // returns all values
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

## License

Copyright (c) 2018-present [Andrey Mak](https://github.com/oceanicdev)

Licensed under [MIT License](./LICENSE)