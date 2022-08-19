package main

import "github.com/semanticist21/go-pv-simulator/simulation"

func main() {
	simulation.RunSimulation(3, 100, "test")

	// endless loop
	for {
	}
}
