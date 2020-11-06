package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./proverb.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write([]byte("don't communicate by sharing memory share memory by communicating"))
	}
	f, err = os.Open("proverb.txt")
	defer f.Close()
	var slice=make([]byte,1024)
	_,err=f.Read(slice)
	fmt.Println(string(slice))

}