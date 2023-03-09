package main

import "fmt"

func main8() {
	start()
}

func start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in start()")
		}
	}()

	fmt.Println("Called start()")
	part2(0)
}

func part2(i int) {
	if i > 0 {
		fmt.Println("Panicking in part2()")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in part2()")
	fmt.Println("Called part2()")

	part2(1)
}
