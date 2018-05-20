package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i<1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hellow %d\n",i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
