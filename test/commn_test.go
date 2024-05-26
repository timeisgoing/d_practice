package test

import (
	"fmt"
	"testing"
)

func TestTimeConsumin1g(t *testing.T) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("===================5====================")
			//continue  //这一次 后面的不执行了
			//break  //终止这一次循环 ，循环剩余的都不执行了，但是循环外面的还执行
			return //整个方法，到此为止了
		}
		fmt.Println(i)

	}
	fmt.Println("over ")
}
