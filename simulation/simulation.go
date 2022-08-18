package simulation

import (
	"math"
	"math/rand"
	"time"

	"github.com/semanticist21/go-pv-simulator/model"
)

func RunSimulation(secInterval int) *time.Ticker {
	ticker := time.NewTicker(time.Second * time.Duration(secInterval))
	go func() {
		for t := range ticker.C {
			baseTemp := genSimulatedTemp()
			baseHz := genSimulatedHz()

			pvOne := getPvWithData(1, baseTemp, baseHz, getSimulatedGen(t.Hour(), t.Minute()))
			pvTwo := getPvWithData(2, baseTemp, baseHz, getSimulatedGen(t.Hour(), t.Minute()))

			jsonA, _ := pvOne.MarshalJson()
			jsonB, _ := pvTwo.MarshalJson()
		}
	}()

	return ticker
}

func getPvWithData(id int, baseTemp float64, baseHz float64, gen float64) *model.Pv {
	pvId := id
	pvGen := gen
	pvHz := genSimulatedHz()
	pvTemp := addSomeMinorTempDifference(baseTemp)
	pvModuleTemp := addSomeMinorTempDifference(pvTemp)

	newPv := &model.Pv{pvId, pvGen, pvHz, pvTemp, pvModuleTemp}

	return newPv
}

// range 0~150kw
func getSimulatedGen(hour int, min int) float64 {
	hours := float64(hour) + float64(min/60)
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

func genSimulatedTemp() float64 {
	return 20 + 10*rand.Float64()
}

func addSomeMinorTempDifference(num float64) float64 {
	return num + 5*(rand.Float64()-0.5)
}
