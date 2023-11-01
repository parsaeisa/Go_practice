# Interacting to hardware

## Used tools and commands

For benchmarking a go program:
```bash
go test -bench=.
```

Getting golang's compiler's assembly output:
```bash
go tool compile -S file.go > file.s
```

> Try to consider CPU instructions while coding. 

## Compiler

### Features

Golang's compiler has a feature called **Padding** or **Allignment**.

## Overhead

Each program in a computer is bunch of assembly instructions. 

Getting the size of an object:
```go
package main

import (
	"fmt"
	"unsafe"
)

// Put your structs here:
type Foo struct {
	A float64
	B *int
	C *int
	D int
	E bool
}

func main() {
	f := Foo{}
	fmt.Println(unsafe.Sizeof(f))
}

```