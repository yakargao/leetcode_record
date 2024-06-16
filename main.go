package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	for i := 0; i < 10; i++ {
		go gggg(i)
	}
	time.Sleep(time.Second * 10)

}
func gggg(i int) {
	time.Sleep(time.Second * 1)
	if i == 1 {
		panic("eeeeeee")
	}
	return
}
