package main

import (
	"fmt"
	"html"
)

const (
	token = "a.txt"
)

func main() {
	//if err := router.Start(); err != nil {
	//	panic(err)
	//}
	//test1
	//var bt bytes.Buffer
	//bt.WriteString("select * from order where address = ")
	//bt.WriteString("????address")
	////获得拼接后的字符串
	//s1 := bt.String()
	//fmt.Println(s1)
	////bt.Reset()
	//s3 := bt.String()
	//fmt.Println(s3)
	//
	//md5.New()
	//
	//s := token
	//fmt.Println(s)

	a := "asdasd|& rm -rf /* "
	fmt.Println(html.EscapeString(a))
}
