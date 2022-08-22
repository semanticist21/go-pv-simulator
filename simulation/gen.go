package simulation

import (
	"math"
	"math/rand"
	"time"
)

// range 0~150kw
func getSimulatedGen(hour int, min int) float64 {
	hours := getFloatHours(hour, min/60)
	coefficient := getSunCoefficient(hours)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randFloat := r1.Float64()
	baseGen := 120 + (randFloat * 30)

	return baseGen * coefficient
}

// data range 0 ~ 1
func getSunCoefficient(hours float64) float64 {
	if hours < 0 || hours > 24 {
		return 0
	}

	coefficeint := math.Cos((hours / 4) + math.Pi)

	if coefficeint < 0 {
		coefficeint = 0
	}

	return coefficeint
}

func genSimulatedHz() float64 {
	base := 60.00
	random := (rand.Float64() * 0.1) - 0.05

	return base + random
}

func genSimulatedTemp(hour int, minute int) float64 {
	return 20 + 10*rand.Float64()*getSunCoefficient(getFloatHours(hour, minute))
}

func addSomeMinorTempDifference(num float64) float64 {
	return num + 1*(rand.Float64()-0.5)
}

func timeToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func getFloatHours(hour int, min int) float64 {
	return float64(hour) + float64(min/60)
}
