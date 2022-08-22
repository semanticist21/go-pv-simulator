package main

import (
	"fmt"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/server"
	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {
	var targetUrl *string = comm.Make("localhost:8080")
	fmt.Printf("target url is %s", *targetUrl)

	server.StartTestServer(targetUrl)
	simulation.RunSimulation(3, 100, "test", targetUrl)

	// endless loop
	for {
	}
}
