package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/semanticist21/go-pv-simulator/model"
)

var baseUrl string = "192.168.1.166:8080"

func StartTestServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/users/{id}/data", DataRequestHandler)
	http.Handle("/", r)

	go http.ListenAndServe(baseUrl, nil)

	fmt.Println("Server has been Initialized.")
}

func SendPvData(dataPkg *model.DataPkg) {
	// url := fmt.Sprintf("%s/users/%d/data?token=%s", baseUrl, dataPkg.UserId, dataPkg.Token)
	// buff := bytes.NewBuffer(dataPkg.JsonData)

	// resp, err := http.Post(url, "application/json", buff)

	// if err != nil {
	// 	fmt.Println("error :" + err.Error())
	// 	return
	// }

	// fmt.Printf("resp.Status: %v\n", resp.Status)
}

func defaultHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Server is active !!"))
}

func DataRequestHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

	}
	// fmt.Println(r.URL.Path)
	// fmt.Println(r.URL.RawQuery)
	// fmt.Println(r.URL.Path)
	// fmt.Println(r.URL.Query()) // url values
	// fmt.Println(reflect.TypeOf(r.URL.Query()))

	// token=test
	// /users/3/data
	// map[token:[test]]
	// url.Values
}
