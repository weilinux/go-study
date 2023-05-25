package main

import "fmt"

func myExecuteFunc0(s string) {
	fmt.Println("my ex func", s)
}

// decorator pattern: make a function accept and do whatever you wanted to do
func main() {
	Execute0(myExecuteFunc0)

}

// ExecuteFn0 ExecuteFn this is coming from a third-party library
type ExecuteFn0 func(string)

func Execute0(fn ExecuteFn0) {
	fn("foo bar baz")
}
