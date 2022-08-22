package simulation

import (
	"fmt"
	"time"

	"github.com/semanticist21/go-pv-simulator/model"
)

func RunSimulation(secInterval int, userId int, password string, targetUrl *string) {
	ticker := time.NewTicker(time.Second * time.Duration(secInterval))

	fmt.Println("Sending data has been started.")

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

			dataA := &model.DataPkg{UserId: userId, Token: password, JsonData: jsonA}
			dataB := &model.DataPkg{UserId: userId, Token: password, JsonData: jsonB}

			go SendPvData(dataA, *targetUrl)
			go SendPvData(dataB, *targetUrl)
		}
	}()
}

