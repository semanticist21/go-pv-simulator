package simulation

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/semanticist21/go-pv-simulator/comm"
	"github.com/semanticist21/go-pv-simulator/model"
)

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
		fmt.Printf("Sent to %s !!\n", url)
	}
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
