package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/onosproject/aether-graphql/internal/config"
	"github.com/onosproject/aether-graphql/internal/routes"
	"github.com/onosproject/aether-graphql/internal/store"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

var log = logging.GetLogger("main")

func main() {
	cfg := config.New()
	flag.IntVar(&cfg.Port, "port", 8080, "API server port")
	flag.StringVar(&cfg.Onos, "onos-address", "localhost:5150", "Onos Config address")
	flag.StringVar(&cfg.Target, "target", "connectivity-service-v2", "Connectivity Service")
	flag.Parse()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	err := store.ConnectGNMI(cfg.Onos)
	if err != nil {
		log.Fatalf("error connecting to gnmi server: %s\n", err)
	}

	store.SetTarget(cfg.Target)

	r := routes.Run()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Infof("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Infof("Server exiting")
}
