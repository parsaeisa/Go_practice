package main

import "fmt"

func test_panic() {
	fmt.Println("function started")
	panic(fmt.Sprintf("%v", 5))
	fmt.Println("yyy")
	defer fmt.Println("some shit has been defered")
	fmt.Println("zzz")
}

func main() {

	//defers the execution of a function until the surrounding function returns.s
	//defer function calls are pushed onto a stack .
	// last in first out
	// next to defer , we have panic and recover
	defer fmt.Println("function has been ended")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	sum := 0
	for i := 0; i < 10; i++ { // The init and post statements are optional.
		sum += i
	}
	fmt.Println(sum)

	sum = 1

	for sum < 1000 { // while loop
		sum += sum
	}

	fmt.Println(sum)

	// test_panic()

}
