# Concurrency . 

There is some neat tools in golang for parallel programming.

* go func
* wait group
* select
* channels


> Map and array are not thread safe. We cannot edit a single array in multiple seperated threads.

Also stringbuilder library is not thread safe.

To edit a shared variable in multiple threads we must do one of the either way : 
* Semaphore, lock
* Atomic
## Channels

When you are done with a channel you should close it using the `close` function.

When you are waiting on a channel, it can return a value that illustrates the status of channel (closed or open), like this : 
```go
var open bool

if err, open = <-nameErr; open {
    return
}
```

## Wait group

You can use wait group as a seperate tool. You can use it when there is no goroutine.

You can add it and defer it's done in several methods that you want to perform something after their completion.

```go
wg *sync.WaitGroup

// Consider the method f is being called in many other goroutines
func f(){
    wg.Add(1)
    defer wg.Done()
}

n.wg.Wait()
// The operations that you want to perform
```

