# Techniques

## Structure

Golang offers a **flat** project structure. It doesn't offer very hierarchical packaging like Java. 

## Parent type 
consider an interface called `validator.ValidationErrors` implements the error interface ( the error in golang is an interface and by validator I mean [this]("github.com/go-playground/validator/v10")) . Lot of methods return error in their output, but sometimes we need other methods that are defined in `validator.ValidationErrors` but not implemented in `error`.

In such cases by the code below, we can access those methods:
```go
err.(validator.ValidationErrors)
```

## Set

When you need a set (where each elements differ), you can easily use map. 

```go
yourMap = make(map[int]struct{})
yourMap[firstElem] = struct{}{}
yourMap[secondElem] = struct{}{}
```

Then for checking an elements existance, you can use this:

```go
val, ok := yourMap[searchedElem]
```

The `ok` shows that whether your element exist or not. 

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