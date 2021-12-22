package main

import (
	"context"
	"fmt"

	stdlog "log"

	"github.com/saint-yellow/go-pl/toys/distributed-system/business"
	"github.com/saint-yellow/go-pl/toys/distributed-system/log"
	"github.com/saint-yellow/go-pl/toys/distributed-system/registry"
	"github.com/saint-yellow/go-pl/toys/distributed-system/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.BusinessService,
		ServiceURL: serviceAddress,
		RequiredServices: []registry.ServiceName{
			registry.LogService,
		},
		ServiceUpdateURL: serviceAddress + "/service",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(), host, port, r, business.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		fmt.Printf("logging service not found at %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down the business service...")
}