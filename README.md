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

## Topics
- [Data structures](https://github.com/parsaeisa/Go_practice/blob/main/Notes/DataStructure.md)
- [Reflection](https://github.com/parsaeisa/Go_practice/blob/main/Notes/Reflection.md)
- [Techniques](https://github.com/parsaeisa/Go_practice/blob/main/Notes/Techniques.md)
- [Concurrency](https://github.com/parsaeisa/Go_practice/blob/main/concurrency/Concurrency.md)
- [Bash commands](https://github.com/parsaeisa/Go_practice/blob/main/bash.md)
- [Good libraries](https://github.com/parsaeisa/Go_practice/tree/main/Good%20libraries)
- [Built-in libraries](https://github.com/parsaeisa/Go_practice/tree/main/Built-in.md)
- [Interacting with hardware](https://github.com/parsaeisa/Go_practice/blob/main/interacting-to-hardware.md)
- [Environment variables](https://github.com/parsaeisa/Go_practice/blob/main/variables.md)

## Design patterns

Golang implementation on design patterns  : https://refactoring.guru/design-patterns/go

## A brilliant way to learn new things 

**Reading source codes of big projects** is a very simple thing that can always provide us with new things to learn.

> Try to read **one** golang project from github each **week**.

Github repo queue:
- https://github.com/yasharne/go.d.plugin
- https://github.com/mhrlife/bigcache
- https://github.com/leandromoreira/cdn-up-and-running

## To Learn : 
* Goproxy ( try to run one goproxy )
* Go embed
* Go scheduler
* concurrency patterns 
* what was that fast json unmarshalling golang package ? 
* gosec
* Search "advanced golang programming" in google.

## Concepts

### Memory leakage

Memoray leakage happens when the program allocates a memory and cannot ( or forgets ) to de-allocate it. 

Here, for programming languages, there is a feature called garbage collector which de-allocates useless allocated parts of memory. 

For memory there are some names that delve in:
- Stack:stack memory is used for storing local variables and function call information. Each Goroutine in Golang has its own stack, typically with a fixed size (e.g., 2MB). Stack memory is managed efficiently by Golang runtime and is automatically reclaimed when a Goroutine exits.

- Heap: managing heap has more overhead for the code. Heap is used for dynamically allocated data, such as objects created with 'new' or 'make' operators.

There are some metrics which shows the situation of memory leakage: 
- rates of allocation  
- number of live objects
- memory in stack
- memory in heap
- GC duration quantile

Monitoring these metrics in a long time can give a good sight about memory leakage. Memory leakage in a long-time can be very harmful. 