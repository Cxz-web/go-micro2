package main

import "fmt"

type Human struct {
	Name string `json:"name"` // 姓名
	//Gender string `json:"s"`    // 性别，性别的tag表明在json中为s字段
	Age int `json:"Age"` // 年龄
}

func test() (t string) {
	t = "123"
	return t
}

func main() {
	fmt.Println(test())
}
