package main

import (
	"fmt"
	"os"

	"github.com/takama/daemon"
)

func main() {

	service, err := daemon.New("takumakumed", "takumakumemon desu")
	usage := "Usage: takumakumed install | remove | start | stop | status"

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			service.Install()
		case "remove":
			service.Remove()
		case "start":
			service.Start()
		case "stop":
			service.Stop()
		case "status":
			service.Status()
		default:
			fmt.Println(usage)
		}
	}

}
