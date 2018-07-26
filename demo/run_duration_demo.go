package main

import (
	"time"
	"os"
	"fmt"
	"github.com/petelin/restart"
)

func main() {
	fmt.Println("i am ", os.Getpid())
	go func() {
		<-time.After(time.Second * 5)
		os.Exit(0)
	}()
	restart.RunWithDuration(func() {
		fmt.Println("my father is ", os.Getppid())
		time.Sleep(time.Second * 2)
	}, time.Second)
}
