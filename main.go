package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/takama/daemon"
)

const (
	// name of the service
	name        = "takumakumed"
	description = "daemon of takumakume"
)

type Service struct {
	daemon.Daemon
}

func (service *Service) Runner() (string, error) {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return fmt.Sprintf("Usage: %v [install | remove | start | stop | status]", os.Args[0]), nil
		}
	}

	for {
		Run()
	}

	return "", nil
}

func Run() {
	fmt.Println("hoge")
	time.Sleep(3 * time.Second)
}

func main() {
	d, err = daemon.New(name, description)
	if err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}

	service := &Service{d}
	status, err := service.Runner()
	if err != nil {
		log.Println(status, "\nError:", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
