package main

import (
	"github.com/semanticist21/go-pv-simulator/server"
	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {
	server.StartTestServer()
	simulation.RunSimulation(3, 100, "test")

	// endless loop
	for {
	}
}
