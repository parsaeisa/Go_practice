# Variables

- `GOMAXPROCS`: This variable limits the number of OS's threads which are dedicated to user processes.

The usage of `GOMAXPROCS` variable becomse critical while dealing with throtling. The closer `GOMAXPROCS` is to the **CPU limit**, the less throttling occurs. This fact is described in the [this](https://medium.com/@sharyash81/solving-cpu-throttling-issue-in-golang-applications-before-hitting-the-cpu-limit-in-kubernetes-7d8f40da6477) link.

- `GODEBUG`:

- `GORACE`: Plays the role of **Race detector** in programs which are executed using `-race` flag.