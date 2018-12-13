package main

import (
	"./real"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}
func main() {
	var r Retriever
	r = real.Retriever{UserAgent:"Mozilla/5.0",TimeOut:time.Minute}
	fmt.Printf("%T %v\n",r,r)
	//fmt.Println(download(r))

	//Type assertion
	realRetriever := r.(real.Retriever)
	fmt.Println(realRetriever.TimeOut)


}
