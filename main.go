package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":3030", nil)
	fmt.Println(GetSunCoefficient(11.11))
	fmt.Println(GetSunCoefficient(12.11))
	fmt.Println(GetSunCoefficient(14.11))
	fmt.Println(GetSunCoefficient(16.11))
	fmt.Println(GetSunCoefficient(21.11))
}
