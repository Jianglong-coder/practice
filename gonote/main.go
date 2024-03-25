package main

import (
	"fmt"
	"time"
)

func main() {

	// note.SayHelloWorld()
	// note.EscapedCharacers()
	// note.VariablesAndConstant()
	// fmt.Println(note.Version) //跨包调用的常量
	// note.BasicDateTypes()

	var intChan chan int = make(chan int)
	select {
	case <-intChan:
		fmt.Println("111")
	case t := <-time.After(time.Second):
		fmt.Print(t)
	}
}
