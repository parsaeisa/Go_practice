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
- [Data structures](https://github.com/parsaeisa/Go_practice/blob/main/Notes/DataStructure.md)
- [Reflection](https://github.com/parsaeisa/Go_practice/blob/main/Notes/Reflection.md)
- [Techniques](https://github.com/parsaeisa/Go_practice#Techniques)
- [Concurrency](https://github.com/parsaeisa/Go_practice/blob/main/concurrency/Concurrency.md)
- [Bash commands](https://github.com/parsaeisa/Go_practice/blob/main/bash.md)

### To Learn : 
* Goproxy ( try to run one goproxy )
* Go embed
* Go scheduler
* concurrency patterns 
* what was that fast json unmarshalling golang package ? 
* gosec
* Search "advanced golang programming" in google.

## Techniques

### Parent type 
consider an interface called `validator.ValidationErrors` implements the error interface ( the error in golang is an interface and by validator I mean [this]("github.com/go-playground/validator/v10")) . Lot of methods return error in their output, but sometimes we need other methods that are defined in `validator.ValidationErrors` but not implemented in `error`.

In such cases by the code below, we can access those methods:
```go
err.(validator.ValidationErrors)
```

### Opts
In golang you can put just one optional argument in a method at most.

When you have a method which got more than one optional paramateres, you can use this pattern.

```go
type Opts struct {
	DeliverNew  bool
	DeliverAll  bool
    Middlewares []SubscribeMiddleware
}

type Opt func(s *Opts)

func DeliverNew() Opt {
	return func(s *Opts) {
		s.DeliverNew = true
	}
}

func DeliverAll() Opt {
	return func(s *Opts) {
		s.DeliverAll = true
	}
}

func WithMiddleware(middleware Middleware) Opt {
	return func(s *Opts) {
		s.Middlewares = append(s.Middlewares, middleware)
	}
}
```

In this way you are actually abstracting the parameters that you want to be optional as a single type, Then you can make all of them optional parameters. 

While defining the signature of method (with more than one optional argument), you just add "...Opt", And when you call that method, you add the defined methods (like DelivarAll) in code. 
```go
// definig method 
func method(int param1, int param2, ...opts){
    // oprations
}

// calling the method
method(param1, param2, DeliverALL(), DeliverNew())
```

Also there is something that you have to in the method. You should create a new instance of `Opts` object, then apply each of Opt methods on them.

At last you use that created Opts object's properties for your logic.
```go
func method(int param1, int param2, ...opts){
    var options Opts
	for _, o := range opts {
		o(&options)
	}
    
    // main oprations
}
```