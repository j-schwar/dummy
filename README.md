# Dummy

Dummy is a modular and highly customizable library for randomly generating Go 
structures using reflection.

Dummy is **not** a library for randomly populating user defined structures with
data. Instead, it focuses on generating completely new structures (fields and
all) based on some user supplied configuration. The primary use case for this
library is to generate random structures which are then marshalled into some
other data format (e.g., JSON, XML, etc.) in order to test parsers and other 
similar of programs.

## Installation

```
go get github.com/j-schwar/dummy
```

## Usage

To get started quick, the default configuration will allow you to create some
sufficiently random, yet realistic data.

```go
package main

import (
  "encoding/json"
  "fmt"
  "github.com/j-schwar/dummy"
)

func main() {
  // Generate an interface{} containing some random data.
  data := dummy.Generate()
  
  // Marshall to JSON and print.
  b, err := json.Marshal(data)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(b))
}
```
