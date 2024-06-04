package main

import (
	"auth-service/internal/config"
	"fmt"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: Initialize logger
	
	// TODO: Initialize application

	// TODO: Launch gRPC-server of application
}
