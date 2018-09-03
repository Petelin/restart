package main

import (
	"os"
	"os/exec"
	"log"
	"github.com/petelin/restart"
	"context"
	"os/signal"
	"syscall"
)

func run() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT,
		syscall.SIGKILL, syscall.SIGTERM)

	var args = []string{"-c"}
	args = append(args, os.Args[1:len(os.Args)]...)

	cmd := exec.Command("bash",  args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Fatalln("Cannot start subprocess, exit with err:", err)
	}

	subProcessCtx, subProcessCancel := context.WithCancel(context.Background())
	go func() {
		cmd.Wait()
		subProcessCancel()
	}()

	select {
	case sig := <-c:
		log.Println("receive signal", sig)
		err := cmd.Process.Kill()
		if err != nil {
			log.Fatalln("kill subprocess failed, exit")
		}
		os.Exit(0)
	case <-subProcessCtx.Done():
	}
}

func main() {
	restart.Run(run)
}
