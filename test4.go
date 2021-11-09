package main

import (
	"errors"
	"log"
)

// package foo

var ErrCouldNotOpen = errors.New("could not open")

func open()  error{
	return ErrCouldNotOpen //如果客户端需要检测错误，并且您已使用创建了一个简单的错误 errors.New，请使用一个错误变量。
}

func main()  {
	if errors.Is(open(), ErrCouldNotOpen) {
		// handle
		log.Fatal("已知错误")
	} else {
		panic("unknown error")
	}
}
