package main

import "fmt"

func main() {
	Greeting()
	advanceGreeting("Ponni")
}

func Greeting() {
	fmt.Println("Hey there!")
}

func advanceGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
