# Data structures

Golang has arrays and slices . 

array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.
```go
var a [5]int
```

but slice has a dynamic size . 
```go 
s := make([]string, 3)
```
## Map
It can map from any type of variable to any other variable type : 

m = map[interface{}]interface{}

creating a map : 
```go
likes := make(map[string]int)
```

adding a key/value to map : 
```go
likes["parsa's post 1"] = 1
```

deleting a key/value from map : 
```go
delete(likes, "parsa's post 1")
```

### Variable assignment
It has two types of variable assignment :
- one variable
- two variable

That means you can type 
```go
i := likes["parsa's post 1"]
```
or
```go
i,ok := likes["parsa's post 1"]
```
If the key exists in map, the `ok` variable is true, otherwise it is false.

### Concurrency

Map data structure has some challenges in parallel use. [To complete]