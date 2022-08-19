package simulation

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/semanticist21/go-pv-simulator/model"
	"github.com/semanticist21/go-pv-simulator/server"
)

func RunSimulation(secInterval int, userId int, password string) {
	ticker := time.NewTicker(time.Second * time.Duration(secInterval))

	fmt.Println("Simulation started.")

	go func() {
		for t := range ticker.C {
			pvIdOne := 1
			pvIdTwo := 2

			baseTemp := genSimulatedTemp()
			baseHz := genSimulatedHz()

			pvOne := getPvWithData(pvIdOne, baseTemp, baseHz, t)
			pvTwo := getPvWithData(pvIdTwo, baseTemp, baseHz, t)

			jsonA, _ := pvOne.MarshalJson()
			jsonB, _ := pvTwo.MarshalJson()

			fmt.Println(string(jsonA))
			fmt.Println(string(jsonB))

			dataA := &model.DataPkg{UserId: userId, Password: password, JsonData: string(jsonA)}
			dataB := &model.DataPkg{UserId: userId, Password: password, JsonData: string(jsonB)}

			go server.SendPvData(dataA)
			go server.SendPvData(dataB)
		}
	}()
}

func getPvWithData(id int, baseTemp float64, baseHz float64, t time.Time) *model.Pv {

	pvId := id
	pvGen := getSimulatedGen(t.Hour(), t.Minute())
	pvHz := genSimulatedHz()
	pvTemp := addSomeMinorTempDifference(baseTemp)
	pvModuleTemp := addSomeMinorTempDifference(pvTemp)

	newPv := &model.Pv{Id: pvId, GenkW: pvGen, Hz: pvHz, Temp: pvTemp, ModuleTemp: pvModuleTemp, Time: timeToRFC3339(t)}

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

func timeToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}
