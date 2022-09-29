package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/semanticist21/go-pv-simulator/simulation"
)

func main() {
	// address
	sc := bufio.NewScanner(os.Stdin)

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

	fmt.Println("Real-time? -> Y, ")

	sc.Scan()
	answer := getTrimmed(sc.Text())

	if answer == "Y" {
		genRealtime(sc, &addr)
	} else {
		batchSimulatedData(sc, &addr)
	}

	// endless loop
	for {
	}
}

func batchSimulatedData(sc *bufio.Scanner, addr *string) {
	fmt.Println("It will batch data as by 30min.")
	fmt.Println("Prompt interval(min)")
	interval := promptNum(sc)
	cnt := promptPvCount(sc)

	var quantity int
	for {
		fmt.Println("How many data you want to gen for each Pv?")
		sc.Scan()
		input := sc.Text()
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please prompt number")
		}

		quantity = num
		break
	}

	pwdToken := promptToken(sc)

	simulation.BatchData(interval, cnt, quantity, &pwdToken, addr)

}

func genRealtime(sc *bufio.Scanner, addr *string) {

	fmt.Println("It will send data real-time.")

	// interval
	fmt.Println("Prompt interval(sec).")
	interval := promptNum(sc)

	// pv
	cnt := promptPvCount(sc)

	// Deleted user Id, no longer used
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

	pwdToken := promptToken(sc)

	// fmt.Println("If test, prompt Y (Will deploy local host server)")
	// sc.Scan()
	// answer := sc.Text()
	// TrimmedAnswer := strings.TrimSpace(strings.ToUpper(answer))

	var targetUrl *string = addr
	var token *string = &pwdToken
	//query parameter

	fmt.Printf("Target url is %s.\n", *targetUrl)
	fmt.Printf("Default user token : %s.\n", *token)
	fmt.Printf("PV Data URL would be %s/data/reg?token=%s\n", *targetUrl, *token)
	// fmt.Printf("User id : %d.\n", userId)
	// fmt.Printf("User Name : %s.\n", *userNm)

	// if TrimmedAnswer == "Y" {
	// 	server.StartTestServer(targetUrl)
	// }

	simulation.RunSimulationRealtime(interval, cnt, token, targetUrl)
}

func getTrimmed(val string) string {
	return strings.TrimSpace(strings.ToUpper(val))
}

func promptToken(sc *bufio.Scanner) string {
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

	return pwdToken
}

func promptNum(sc *bufio.Scanner) int {
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

		if num <= 0 {
			fmt.Println("Please prompt positive number.")
			continue
		}

		interval = num
		break
	}
	return interval
}

func promptPvCount(sc *bufio.Scanner) int {
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

	return cnt
}
