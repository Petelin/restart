package main

import (
	"os"
	"time"
	"fmt"
	"github.com/petelin/restart"
)

func main() {
	fmt.Println("i am ", os.Getpid())
	go func() {
		<-time.After(time.Second)
		os.Exit(0)
	}()
	restart.Run(func() {
		fmt.Println("my father is ", os.Getppid())
		time.Sleep(time.Second * 2)
	})
}
