
# Design Patterns

In this directory you can see my practices around golang concurrency design patterns .

here's a list of design patterns that I want to practice : 
* fan out

### Notes 

sometimes in order to adjust timings between multiple channel we use a new channel . 

We will not put anything in this new channel , and just by closing it and using "<-channel" 
in another goroutine , we inform that another goroutine that something happened . 


* To see an example go to fan-in design pattern example . 