package main

import (
	"fmt"
	"github.com/litecy/goproxy/internal/config"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/litecy/goproxy/pkg/services"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("err : %s", err)
	}
	if config.Service != nil && config.Service.S != nil {
		Clean(&config.Service.S)
	} else {
		Clean(nil)
	}
}
func Clean(s *services.Service) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("crashed, err: %s\nstack:\n%s", e, string(debug.Stack()))
			}
		}()
		for range signalChan {
			log.Println("Received an interrupt, stopping services...")
			if s != nil && *s != nil {
				(*s).Clean()
			}
			if config.Cmd != nil {
				log.Printf("clean process %d", config.Cmd.Process.Pid)
				config.Cmd.Process.Kill()
			}
			if *config.IsDebug {
				config.SaveProfiling()
			}
			cleanupDone <- true
		}
	}()
	<-cleanupDone
	os.Exit(0)
}
