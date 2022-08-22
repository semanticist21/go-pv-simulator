package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/server"
	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Prompt address to send data. Just Enter for \"localhost:8080\"")

	addr, _ := reader.ReadString('\n')
	if addr == "\r\n" {
		addr = "localhost:8080"
	}

	fmt.Println("If test, prompt Y")
	answer, _ := reader.ReadString('\n')
	UpperedAnswer := strings.ToUpper(answer)

	var targetUrl *string = comm.Make(strings.TrimSpace(addr))
	fmt.Printf("target url is %s\n", *targetUrl)

	if strings.TrimSpace(UpperedAnswer) == "Y" {
		server.StartTestServer(targetUrl)
	}

	simulation.RunSimulation(3, 100, "test", targetUrl)

	// endless loop
	for {
	}
}
