package simulator

import "math"

func BeginSimulation() {

}

func GetSunCoefficient(hour float64) float64 {
	if hour < 0 || hour > 24 {
		return 0
	}

	coefficient := 10 * math.Cos((hour/4)+math.Pi)

	if coefficient < 0 {
		coefficient = 0
	}

	return coefficient
}
