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

### To Learn : 
* Goproxy ( try to run one goproxy )
* Go embed
* Go scheduler
* concurrency patterns 
* what was that fast json unmarshalling golang package ? 
* gosec

## Techniques

### Parent type 
consider an interface called `validator.ValidationErrors` implements the error interface ( the error in golang is an interface and by validator I mean [this]("github.com/go-playground/validator/v10")) . Lot of methods return error in their output, but sometimes we need other methods that are defined in `validator.ValidationErrors` but not implemented in `error`.

In such cases by the code below, we can access those methods:
```go
err.(validator.ValidationErrors)
```
