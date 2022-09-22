package simulation

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/model"
)

func RunSimulation(secInterval int, userId int, userNm *string, targetUrl *string) {
	ticker := time.NewTicker(time.Second * time.Duration(secInterval))

	fmt.Println("Sending data has been started.")

	go func() {
		for t := range ticker.C {
			pvIdOne := 1
			pvIdTwo := 2

			baseTemp := genSimulatedTemp(t.Hour(), t.Minute())
			baseHz := genSimulatedHz()

			pvOne := getPvWithData(pvIdOne, baseTemp, baseHz, t)
			pvTwo := getPvWithData(pvIdTwo, baseTemp, baseHz, t)

			jsonA, _ := pvOne.MarshalJson()
			jsonB, _ := pvTwo.MarshalJson()

			dataA := &model.DataPkg{UserNm: *userNm, JsonData: jsonA}
			dataB := &model.DataPkg{UserNm: *userNm, JsonData: jsonB}

			go SendPvData(dataA, *targetUrl)
			go SendPvData(dataB, *targetUrl)
		}
	}()
}

func getPvWithData(id int, baseTemp float64, baseHz float64, t time.Time) *model.Pv {

	pvId := id
	pvGen := getSimulatedGen(t.Hour(), t.Minute())
	pvHz := genSimulatedHz()
	pvTemp := addSomeMinorTempDifference(baseTemp)
	pvModuleTemp := addSomeMinorTempDifference(pvTemp)

	newPv := &model.Pv{PvId: pvId, GenkW: pvGen, Hz: pvHz, Temp: pvTemp, ModuleTemp: pvModuleTemp, Time: timeToRFC3339(t)}

	return newPv
}

// gen
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
	return 20 + 5*rand.Float64()*getSunCoefficient(getFloatHours(hour, minute))
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

//send data
var protocol *string = comm.Make("http://")

func SendPvData(dataPkg *model.DataPkg, targetUrl string) {
	targetUrlWithProtocol := *protocol + targetUrl
	url := fmt.Sprintf("%s/users/%d/data?token=%s", targetUrlWithProtocol, dataPkg.UserId, dataPkg.Token)

	result, _ := dataPkg.MarshalJson()
	buff := bytes.NewBuffer(result)

	resp, err := http.Post(url, "application/json", buff)

	if err != nil {
		fmt.Println("error :" + err.Error())
		return
	}

	if strings.Compare(resp.Status, "200 OK") == 0 {
		if !strings.Contains(url, "localhost") {
			fmt.Printf("Sent to %s !!\n", url)
		}
	}
}
