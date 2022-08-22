package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/server"
	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("Prompt address to send data. Just Enter for \"localhost:8080\".")
	sc.Scan()
	addr := sc.Text()

	if addr == "" {
		fmt.Println("localhost:8080 was selected.")
		addr = "localhost:8080"
	}

	var interval int

	for {
		fmt.Println("Prompt Interval(sec)")
		sc.Scan()
		input := sc.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please prompt number")
			continue
		}

		interval = num
		break
	}

	fmt.Println("If test, prompt Y")
	sc.Scan()
	answer := sc.Text()
	UpperedAnswer := strings.ToUpper(answer)

	var targetUrl *string = comm.Make(addr)
	fmt.Printf("target url is %s\n", *targetUrl)

	if strings.TrimSpace(UpperedAnswer) == "Y" {
		server.StartTestServer(targetUrl)
	}

	simulation.RunSimulation(interval, 100, "test", targetUrl)

	// endless loop
	for {
	}
}
