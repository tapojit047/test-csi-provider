package main

import (
	fake "github.com/tapojit047/test-csi-provider/server"
	"log"
	"os"
)

func main() {
	// Define the Unix socket path
	socketPath := "/var/lib/csi/test-csi.sock"

	// Clean up any existing socket
	if _, err := os.Stat(socketPath); err == nil {
		if err := os.Remove(socketPath); err != nil {
			log.Fatalf("Failed to remove existing socket: %v", err)
		}
	}

	// Create the mock server
	mockServer, err := fake.NewMocKCSIProviderServer(socketPath)
	if err != nil {
		log.Fatalf("Failed to create MockCSIProviderServer: %v", err)
	}

	// Start the server
	if err := mockServer.Start(); err != nil {
		log.Fatalf("Failed to start MockCSIProviderServer: %v", err)
	}
	defer mockServer.Stop()

	log.Printf("Mock CSI Provider Server started at %s", socketPath)

	// Keep the server running
	select {}
}
