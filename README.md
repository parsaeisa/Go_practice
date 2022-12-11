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
