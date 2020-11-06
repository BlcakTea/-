package main

import (
	"fmt"
	"time"
)


func p1(count chan int )  {
    var i int
	for i=0;i<=100;i++{
       count<-i
        if i%2==0{
        	fmt.Println(i)
		}
	}

}
func p2(count chan int )  {
	var i int
	for i=0;i<=100;i++{
		<-count
		if i%2==1{
			fmt.Println(i)
		}
	}

}
func main()  {
	number:=make(chan int)
	go p1(number)
	go p2(number)
	time.Sleep(time.Second *2)
}