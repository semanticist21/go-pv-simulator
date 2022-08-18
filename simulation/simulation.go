package simulation

import "math"

func RunSimulation() {

}

func GetSunCoefficient(hour float64) float64 {
	if hour < 0 || hour > 24 {
		return 0
	}

	coefficeint := hourCosFunc(hour)

	if coefficeint < 0 {
		coefficeint = 0
	}

	return coefficeint
}

func hourCosFunc(hour float64) float64 {
	return 10 * math.Cos((hour/4)*math.Pi)
}
