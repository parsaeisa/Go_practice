# Go_practice

after programming in many languages like C,C++,C#,java, python and etc ,
I find Golang the best language in software backend developing because of four reasons :

+ Support multithreading programming strongly and easily .
+ Having shortcuts of python . 
+ Start a new project with no unnecessary files ( Its not like laravel !! )
+ Give programmer the most freedom !! 

I started with tour of Go and this repo is the play ground for that tour . you can follow that tour in the link  below : 

https://go.dev/tour/


after this tour I learned fundamentals of programming with golang so that I could run a url shortener backend server .

+ you can see my url shortener in following repo : 
https://github.com/parsaeisa/golang_url_shortener

### Topics
- [Data structures](https://github.com/parsaeisa/Go_practice#data-structures)
- [Reflection](https://github.com/parsaeisa/Go_practice#reflection)

### To Learn : 
* strings.Builder
* goqu : a package which generates sql queries for all of drivers like mysql , mariadb , postgresql , sqlite and etc . 

## Data structures

Golang has arrays and slices . 

array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios.
```go
var a [5]int
```

but slice has a dynamic size . 
```go 
s := make([]string, 3)
```
### Map
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

#### Variable assignment
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

## Reflection

consider we have an interface{}. We can get it's type using the code below:
```go
val := reflect.ValueOf(res)
```

Reflect package has some types in a pre-defined way on it's own. types like : `reflect.Ptr` and `reflect.Struct`.

If an interface{} is Ptr we can indirect it using : 
```go
val = reflect.Indirect(val)
```

`val.Type.String()` : has the type of the object.

### Get attributes
An object can be struct and have some fields. 

`structType.NumField()` returns the number of fields.

`structType.Field(index)` returns the field in a specific index. 

The 'structType' is the output of val.Type() 

Also we can check the tags of each field. For example to get the "yml" tag : 
```go
tag := field.Tag
fieldTags := tag.Get("yaml")
```

The field is the output of the method `Field` that was mentioned before.