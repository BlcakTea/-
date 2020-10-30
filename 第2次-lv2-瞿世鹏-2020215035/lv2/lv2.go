package main

import "fmt"

func Receiver(a interface{}) {
	switch a.(type) {
	case int:fmt.Println("这个是int")
	case bool:fmt.Println("这个是bool")
	case string:fmt.Println("这个是string")
	}
}

func main() {
	Receiver("永远相信美好的事情即将发生")
	Receiver(0406)
	Receiver(true)
}
