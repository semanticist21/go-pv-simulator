package simulation

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

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
