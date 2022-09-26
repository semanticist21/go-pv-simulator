package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {

	fmt.Println("Real-time? -> Y, ")
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	answer := getTrimmed(sc.Text())

	if answer == "Y" {
		genRealtime(sc)
	} else {
		putSimulatedData(sc)
	}

	// endless loop
	for {
	}
}

func putSimulatedData(sc *bufio.Scanner) {
	panic("unimplemented")
}

func genRealtime(sc *bufio.Scanner) {
	// address
	var addr string
	fmt.Println("Prompt full address to send data. Just enter for \"localhost:8080\".")

	for {
		sc.Scan()
		input := sc.Text()

		if !strings.Contains(input, "http") {
			fmt.Println("Include Protocol. https://")
			continue
		}

		if input == "" {
			fmt.Println("http://localhost:8080 was selected.")
			addr = "http://localhost:8080"
		}

		addr = input
		break
	}

	// interval
	var interval int
	fmt.Println("Prompt interval(sec).")

	for {
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

	// pv
	fmt.Println("How many PVs to generate?")
	var cnt int

	for {
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

	// Deleted user Id
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

	var pwdToken string

	for {
		defaultToken := "test"

		fmt.Printf("Prompt auth token. default :: %s\n", defaultToken)
		sc.Scan()
		input := sc.Text()

		if input == "" {
			input = defaultToken
		}

		pwdToken = input
		break
	}

	// fmt.Println("If test, prompt Y (Will deploy local host server)")
	// sc.Scan()
	// answer := sc.Text()
	// TrimmedAnswer := strings.TrimSpace(strings.ToUpper(answer))

	var targetUrl *string = comm.Make(addr)
	var token *string = comm.Make(pwdToken)
	//query parameter

	fmt.Printf("Target url is %s.\n", *targetUrl)
	fmt.Printf("Default user token : %s.\n", *token)
	fmt.Printf("PV Data URL would be %s/data/reg?token=%s\n", *targetUrl, *token)
	// fmt.Printf("User id : %d.\n", userId)
	// fmt.Printf("User Name : %s.\n", *userNm)

	// if TrimmedAnswer == "Y" {
	// 	server.StartTestServer(targetUrl)
	// }

	simulation.RunSimulation(interval, targetUrl, cnt, token)
}

func getTrimmed(val string) string {
	return strings.TrimSpace(strings.ToUpper(val))
}
