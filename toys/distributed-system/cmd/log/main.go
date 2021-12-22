package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/saint-yellow/go-pl/toys/distributed-system/log"
	"github.com/saint-yellow/go-pl/toys/distributed-system/registry"
	"github.com/saint-yellow/go-pl/toys/distributed-system/service"
)

func main() {
	log.Run("./distributed-system.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s",host, port)
	reg := registry.Registration{
		ServiceName: "Log Service",
		ServiceURL: serviceAddress,
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}
	ctx,err := service.Start(context.Background(), host, port, reg, log.RegisterHandlers)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down the log service...")
}