package simulation

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/semanticist21/go-pv-simulator/model"
)

func RunSimulationRealtime(secInterval int, quantity int, token *string, targetUrl *string) {
	ticker := time.NewTicker(time.Second * time.Duration(secInterval))

	fmt.Println("Sending data has been started.")

	go func() {
		for t := range ticker.C {
			for i := 0; i < quantity; i++ {
				pvId := i
				baseTemp := genSimulatedTemp(t.Hour(), t.Minute())
				baseHz := genSimulatedHz()

				pv := getPvWithData(pvId, baseTemp, baseHz, t)
				go SendPvData(pv, targetUrl, token)
			}

			// pvIdOne := 1
			// pvIdTwo := 2

			// baseTemp := genSimulatedTemp(t.Hour(), t.Minute())
			// baseHz := genSimulatedHz()

			// pvOne := getPvWithData(pvIdOne, baseTemp, baseHz, t)
			// pvTwo := getPvWithData(pvIdTwo, baseTemp, baseHz, t)

			// jsonA := pvOne
			// jsonB := pvTwo

			// dataA := &model.DataPkg{UserId: userId, JsonData: *jsonA}
			// dataB := &model.DataPkg{UserId: userId, JsonData: *jsonB}

			// go SendPvData(pvOne, *targetUrl)
			// go SendPvData(pvTwo, *targetUrl)
		}
	}()
}

func BatchData(minInterval int, cnt int, quantity int, token *string, targetUrl *string) {

}

func getPvWithData(id int, baseTemp float64, baseHz float64, t time.Time) *model.Pv {

	pvId := id
	pvGen := getSimulatedGen(t)
	pvHz := genSimulatedHz()
	pvTemp := addSomeMinorTempDifference(baseTemp)
	pvModuleTemp := addSomeMinorTempDifference(pvTemp)

	newPv := &model.Pv{PvId: pvId, GenkW: pvGen, Hz: pvHz, Temp: pvTemp, ModuleTemp: pvModuleTemp, Time: timeToRFC3339(t)}

	return newPv
}

// gen
// range 0~150kw
func getSimulatedGen(t time.Time) float64 {
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()

	hours := getFloatHours(hour, min/60)
	coefficient := getSunCoefficient(hours)

	//cloud
	s0 := rand.NewSource(int64(day))
	r0 := rand.New(s0)

	cloud := r0.Float64()
	if cloud <= 0.4 {
		cloud = 0.4
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	randFloat := r1.Float64()
	baseGen := (120 + (randFloat * 30)) * cloud

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
	return num + 0.2*(rand.Float64()-0.5)
}

func timeToRFC3339(t time.Time) string {
	return t.Format(time.RFC3339)
}

func getFloatHours(hour int, min int) float64 {
	return float64(hour) + float64(min/60)
}

//send data
// var protocol *string = comm.Make("http://")

func SendPvData(dataPkg *model.Pv, targetUrl *string, token *string) {
	// targetUrlWithProtocol := *protocol + targetUrl
	url := fmt.Sprintf("%s/data/reg?token=%s", *targetUrl, *token)

	result, _ := dataPkg.MarshalJson()
	buff := bytes.NewBuffer(result)

	resp, err := http.Post(url, "application/json", buff)

	if err != nil {
		fmt.Println("Error ::" + err.Error())
		return
	}

	if strings.Compare(resp.Status, "200 OK") == 0 {
		if !strings.Contains(url, "localhost") {
			fmt.Printf("Sent to %s !!\n", url)
		}
	} else {
		fmt.Printf("Error :: %s\n", resp.Status)
	}
}
