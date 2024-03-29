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