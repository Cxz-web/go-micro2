package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type pet struct {

}

func (p *pet) speak() {
	fmt.Print(123)
}


type test struct {
	pet
}

func (p *test) speak() {
	fmt.Print(222)
}

func myError() {

	defer func() {
		if s := recover();s !=nil {
			fmt.Println(s)
		}
	}()

	panic("error")
}


func main() {
	myError()
}



func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGUSR2)
	for {
		s := <-c
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
	}
}
