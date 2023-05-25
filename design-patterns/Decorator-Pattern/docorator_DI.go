package main

import "fmt"

// DB 我们不想简单的打印而已，我们想在被调用函数中实现访问DB
type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing into db", value)
	return nil
}

// 现在想在这个被调用中访问DB, 这是就是DI，通过函数注入了DB对象
func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		fmt.Println("my ex func before db", s)
		_ = db.Store(s)
	}
}

// decorator pattern: make a function accept and do whatever you wanted to do
func main() {
	// 现在想在这个被调用中访问DB, 这是就是DI，通过函数注入了DB对象
	// 重点是这个函数接收的是接口而已，实现该接口的类型都可以注入进去
	s := &Store{}
	Execute(myExecuteFunc(s))

}

// ExecuteFn this is coming from a third-party library
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("foo bar baz")
}
