package main

import "fmt"

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() { println("women")}

type S2 struct{}

func (s *S2) f() { println("nimen") }



func main()  {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i.f() //值接收器调用方法
	i = s1Ptr //值接收器方法集是指针接收器方法集的子集,反之不是；比如i既可以调用指针接收器方法，又可以调用值接收器方法
	i.f()  //指针接收器调用方法
	i = s2Ptr
	println(i)
	i=s2Val // S2值接收器没有实现f方法
	fmt.Printf("%#v",s2Val)
}
