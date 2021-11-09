package main

import (
	"fmt"
	"log"
)

// package foo

type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func IsNotFoundError(err error) bool { //直接导出自定义错误类型是错误的，但是由于自定义错误已经成为公共api的一部分，所以我们可以公开匹配器
	_, ok := err.(errNotFound)
	return ok
}

func Open(file string) error {
	return errNotFound{file: file}
}



func main()  {
	// package bar

	if err := Open("foo"); err != nil {
		if IsNotFoundError(err) {
			// handle
			log.Fatal("not found error")
		} else {
			panic("unknown error")
		}
	}
}
