# Usefull libraries

Some handy packages are shown here . 

## Faker

It fills a struct with random values . It is beneficial for unit testing . 
Example  :
```go
package main

import (
	"github.com/bxcodec/faker/v4"
)

type Car struct {
	Number uint32
	Name   string
	Year   uint
}

func main() {

	var car Car
	errFakeData := faker.FakeData(&car)
	if errFakeData != nil {
		panic(errFakeData)
	} else {
		println(car.Number)
		println(car.Name)
		println(car.Year)
	}
}

```

For randomizing an array of values , use the code below : 
```go
var cars []Car
errFakeData := faker.FakeData(&cars, options.WithRandomMapAndSliceMaxSize(uint(5)))
```

You can also add a MinSize in options . 
