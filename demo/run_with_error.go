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
		<-time.After(time.Second * 5)
		os.Exit(0)
	}()
	restart.Run(func() {
		fmt.Printf("i am %d, my father is %d\n", os.Getpid(), os.Getppid())
		time.Sleep(time.Second * 1)
		os.Exit(restart.DirectReturn)
	})
}
