package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"olinker/internal/api"
	"olinker/internal/core"
	"olinker/internal/system"
	"olinker/internal/vendors"
)

func main() {
	os.MkdirAll("logs", 0755)
	logFile, err := os.OpenFile("logs/olinker.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(logFile)
	} else {
		log.Println("WARNING: Could not open logs/olinker.log, falling back to stdout")
	}

	configFile, err := os.ReadFile("configs/config.json")
	if err != nil {
		log.Fatalf("Failed to read config.json: %v", err)
	}

	var config core.VendorConfig
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	log.Printf("Starting oLinker, chosen vendor: %s", config.Vendor)

	vendor, err := vendors.LoadVendor(config)
	if err != nil {
		log.Fatalf("Failed to initialize vendor SDK: %v", err)
	}

	jobQueue := core.NewJobQueue()
	jobCtx, cancelJobs := context.WithCancel(context.Background())
	defer cancelJobs()
	jobQueue.Start(jobCtx)

	encodeService := core.NewEncodeService(jobQueue, vendor)
	httpServer := api.NewServer(config.Port, encodeService)

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	url := fmt.Sprintf("http://localhost:%d", config.Port)
	log.Printf("Opening browser at %s", url)
	if err := system.OpenBrowser(url); err != nil {
		log.Printf("Failed to open browser: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down oLinker...")
}
