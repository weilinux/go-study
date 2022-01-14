package main

import "fmt"

type i1 interface {
	start(string)
	stop(string)
}

type s1 struct {
	s11 string
}

func (s s1) start(ss string) {
	fmt.Printf("s1 start方法：%v\n", ss)
}

func (s s1) stop(ss string) string {
	return fmt.Sprintf("s1 stop方法：%v\n", ss)
}

type s2 struct {
	s22 string
}

func (s s2) start(ss string) {
	fmt.Printf("s2 start方法：%v\n", ss)
}

func (s s2) stop(ss string) {
	fmt.Printf("s2 stop方法：%v\n", ss)
}

type s3 struct {
}

func (s s3) ope(i i1, ss string) {
	i.start(ss)
	i.stop(ss)
}

func main() {
	s11 := s1{"s11"}
	fmt.Printf("s11 t: %T v: %v\n", s11, s11)
	s11.start("s11")
	fmt.Printf("s11 t: %T v: %v\n", s11.stop("s11"), s11.stop("s11"))

	s22 := s2{"s22"}
	fmt.Printf("s22 t: %T v: %v\n", s22, s22)

	/*var i11 i1 = s11
	fmt.Printf("i11 t: %T v: %v\n", i11, i11)*/

	var i22 i1 = s22
	fmt.Printf("i22 t: %T v: %v\n", i22, i22)

	s33 := s3{}
	fmt.Printf("i33 t: %T v: %v\n", s33, s33)

	//s33.ope(i11, "s11")
}
