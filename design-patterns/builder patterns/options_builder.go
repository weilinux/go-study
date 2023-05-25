package main

import "fmt"

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 100,
		id:      "default",
		tls:     false,
	}
}

func withTLS(opt *Opts) {
	opt.tls = true
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func withMaxConn(n int) OptFunc {
	return func(opt *Opts) {
		opt.maxConn = n
	}
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func main() {
	s := newServer(withTLS, withMaxConn(200), withID("weilin"))
	fmt.Println(s)
}
