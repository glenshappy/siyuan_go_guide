package main

import (
	"fmt"
	"log"
)

type errNotFound struct { //如果您有可能需要客户端检测的错误，并且想向其中添加更多信息（例如，它不是静态字符串），则应使用自定义类型。
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use() {
	if err := open("testfile.txt"); err != nil {
		if _, ok := err.(errNotFound); ok {
			// handle
			log.Fatal("not found error")
		} else {
			panic("unknown error")
		}
	}
}

func main()  {
	use()
}
