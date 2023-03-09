# Reflection

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

## Get attributes
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