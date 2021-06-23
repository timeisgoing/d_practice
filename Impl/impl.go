package Impl

import (
	"io/ioutil"
	"net/http"
)

type Retriever struct {
}

func (Retriever) Get(msg string) string {
	get, err := http.Get(msg)
	if err != nil {
		panic(err)
	}
	defer get.Body.Close() //defer是什么意思
	all, _ := ioutil.ReadAll(get.Body)
	return string(all)
}
