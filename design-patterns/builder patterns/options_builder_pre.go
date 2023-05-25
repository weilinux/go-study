package main

import "fmt"

// options pattern 或者 build pattern
// server, err := mypkg.MakeServer("blah").Http(8080).MaxClients(100).Done()

type Server0 struct {
	maxConn int
	id      string
	tls     bool
}

func newServer0(maxConn int, id string, tls bool) *Server0 {
	return &Server0{
		maxConn: maxConn,
		id:      id,
		tls:     tls,
	}
}

func main() {
	// 你会发现下面这个初始化的方案没有灵活性，不能灵活地控制参数写入对象
	s := newServer0(1, "foo", false)
	fmt.Println(s)
}
