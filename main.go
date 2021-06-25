package main

import (
	"awesomeProject/Impl"
	"awesomeProject/Impl2"
	"fmt"
)
//各自对象调用各自的方法
func Getretriever() Impl.Retriever {
	return Impl.Retriever{} //Impl使用的是impl.go的package名字,如何引用别的文件中的结构
}
func Getretriever2() Impl2.Retriever {
	return Impl2.Retriever{} //Impl使用的是impl.go的package名字,如何引用别的文件中的结构
}

func main() {
/*	//各自对象调用各自的方法
	var retriever Impl.Retriever = Getretriever() //用var声明的对象都是用=赋值;没有var才可以使用:=赋值
	get := retriever.Get("http://www.baidu.com")
	fmt.Println(get)

	var retriever2 Impl2.Retriever = Getretriever2()
	get2 := retriever2.Get("http://www.baidu.com")
	fmt.Println(get2)*/
	fmt.Println("====================================")
	var r xxoo
	r = Impl.Retriever{}
	if is,err := r.(Impl2.Retriever);err{
		fmt.Println(is)
	}else {
		fmt.Println("报错了，不是这个类型的")
	}

}
//===============================================================

//使用接口解决调用该参数问题泛型化
type xxoo interface {
	Get(string) string
}

func main1()  {
	var  r xxoo =Getretriever()    //可以用xxoo类型接收retriever类型的结构，可见接口中的类型就是个泛型
	get := r.Get("http://www.baidu.com")
	fmt.Println(get)

	var  r2 xxoo =Getretriever2()   //这俩都是retriever类型的结构,但是调用Get()方法时,去而根据他们的真是类型去执行
	get2 := r2.Get("http://www.baidu.com")
	fmt.Println(get2)

}
//总结
//1.结构叫什么不重要,重要的是要有和真实实现相同的方法, 当然接口中的方法是假的,没有实现
//2.用接口的类型接收各个真实实现,看上去就是java那种父类和子类的效果了