package main

import (
	"awesomeProject/retriever/mock"
	"fmt"
)

//调用者这里指定接口,
//接口中实现那些方法,就是接口的定义了
//实现者呢,实现时隐式的
type Retriever interface {
	Gets(string) string
}

func download(r Retriever) string {
	return r.Gets("http://www.baidu.com")
}
func main() {
	var r Retriever =mock.Retriever{"mock" }
	//var r2 Retriever =real.Retriever{}

	fmt.Println(download(r))

}
