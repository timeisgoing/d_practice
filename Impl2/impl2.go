package Impl2

type Retriever struct {
}

func (Retriever) Get(msg string) string {

	return string(" 这是impl2的Get()")
}
