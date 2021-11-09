package main

import "fmt"

type Student struct {
	name string
}

func main()  {
	var s Student
	var s1 Student
	s=Student{}
	println(s==s1)
	fmt.Printf("%#v",s)
}
