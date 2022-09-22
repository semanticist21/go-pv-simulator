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

	fmt.Println("Prompt full address to send data. Just enter for \"localhost:8080\".")
	fmt.Println("Include Protocol. https://")
	sc.Scan()
	addr := sc.Text()

	if addr == "" {
		fmt.Println("http://localhost:8080 was selected.")
		addr = "http://localhost:8080"
	}

	var interval int

	for {
		fmt.Println("Prompt interval(sec).")
		sc.Scan()
		input := sc.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please prompt number.")
			continue
		}

		interval = num
		break
	}

	var cnt int

	for {
		fmt.Println("Prompt number of pv.")
		sc.Scan()
		input := sc.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please prompt number.")
			continue
		}

		if num <= 0 {
			fmt.Println("Should not be less than 1.")
			continue
		}

		cnt = num
		break
	}

	// var userId int

	// for {
	// 	fmt.Println("Prompt user id.")
	// 	sc.Scan()
	// 	input := sc.Text()
	// 	num, err := strconv.Atoi(input)

	// 	if err != nil {
	// 		fmt.Println("Please prompt number.")
	// 		continue
	// 	}

	// 	userId = num
	// 	break
	// }

	// var userNm *string

	// for {
	// 	fmt.Println("Prompt name.")
	// 	sc.Scan()
	// 	input := sc.Text()

	// 	if reflect.String != reflect.TypeOf(input).Kind() {
	// 		fmt.Println("Please prompt number.")
	// 		continue
	// 	}

	// 	userNm = comm.Make(input)
	// 	break
	// }

	// var token string

	// for {
	// 	fmt.Println("Prompt auth token. default :: test")
	// 	sc.Scan()
	// 	input := sc.Text()

	// 	t := reflect.TypeOf(input).Kind()

	// 	if t != reflect.String {
	// 		fmt.Println("Prompt auth token.")
	// 		continue
	// 	}

	// 	token = input
	// 	break
	// }

	fmt.Println("If test, prompt Y (Will deploy local host server)")
	sc.Scan()
	answer := sc.Text()
	UpperedAnswer := strings.ToUpper(answer)

	var targetUrl *string = comm.Make(addr)
	var token *string = comm.Make("swDev")
	//query parameter

	fmt.Printf("Target url is %s.\n", *targetUrl)
	// fmt.Printf("User id : %d.\n", userId)
	// fmt.Printf("User Name : %s.\n", *userNm)
	// fmt.Printf("Default user token : %s.\n", token)
	// fmt.Printf("User Reg URL would be %s/user/reg\n", *targetUrl)
	fmt.Printf("PV Data URL would be %s/data/reg?token=%s\n", *targetUrl, *token)

	if strings.TrimSpace(UpperedAnswer) == "Y" {
		server.StartTestServer(targetUrl)
	}

	// simulation.RunSimulation(interval, userId, token, targetUrl)
	simulation.RunSimulation(interval, targetUrl, cnt, token)

	// endless loop
	for {
	}
}
