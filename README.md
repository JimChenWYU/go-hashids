# go-hashids [![Build Status](https://travis-ci.org/JimChenWYU/go-hashids.svg?branch=master)](https://travis-ci.org/JimChenWYU/go-hashids) [![GoDoc](https://godoc.org/github.com/JimChenWYU/go-hashids?status.svg)](https://godoc.org/github.com/JimChenWYU/go-hashids)

---

Golang implementation of http://www.hashids.org

### Installing

```bash
$ go get -u github.com/jimchenwyu/go-hashids
```

### Usage

```go
package main

import (
    "fmt"
    "github.com/jimchenwyu/go-hashids"
)

func main() {
    salt := "foobar"
    minLen := 8
    alphabet := hashids.DEFAULT_ALPHABET
    config := hashids.NewHashidConfig(salt, minLen, alphabet)
    h := hashids.NewHashidsObject(config)
    
    fmt.Println(h.Encode([]int{1, 2, 3})) // oGvuDTyl
}
```

### License

[MIT](LICENSE)
