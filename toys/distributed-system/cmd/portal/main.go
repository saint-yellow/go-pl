package main

import (
	"context"
	"fmt"
	stdlog "log"

	"github.com/saint-yellow/go-pl/toys/distributed-system/log"
	"github.com/saint-yellow/go-pl/toys/distributed-system/portal"
	"github.com/saint-yellow/go-pl/toys/distributed-system/registry"
	"github.com/saint-yellow/go-pl/toys/distributed-system/service"
)

func main() {
	err := portal.ImportTemplates()
	if err != nil {
		stdlog.Fatal(err)
	}

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.PortalApplication,
		ServiceURL: serviceAddress,
		RequiredServices: []registry.ServiceName{
			registry.LogService,
			registry.BusinessService,
		},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL: serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(), host, port, r, portal.RegisterHandlers)
	if err != nil {
		stdlog.Fatal(err)
	}
	
	if logProvider, err := registry.GetProvider(registry.LogService); err != nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("shutting down the portal application...")
}