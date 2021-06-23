package mock

type  Retriever struct {
	Contents string
}

func (r Retriever) Gets(s string) string {
	return "mock"
}

//看见左边的箭头了吗
//idea已经识别出来我们实现了接口,
//大胆推测go的解耦实现时判断 struct中的方法和interface是不是同名
//是,就是实现了这个接口了

