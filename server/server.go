package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/semanticist21/go-pv-simulator/model"
)

func StartTestServer(targetUrl *string) {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.HandleFunc("/users/{id}/data", DataRequestHandler)
	http.Handle("/", r)

	go http.ListenAndServe(*targetUrl, nil)

	fmt.Println("Server has been Initialized.")
}

func defaultHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Server is active !!"))
}

func DataRequestHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		data := r.URL.Query()
		token := data["token"][0]

		if token != "test" {
			fmt.Printf("token is not correct :: %s was given.\n", token)
		}

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		dataPkg := new(model.DataPkg)
		dataPkg.UnMarshalJson(reqBody)

		pvData := new(model.Pv)
		pvData.UnMarshalJson(dataPkg.JsonData)

		// fmt.Println(dataPkg.UserId)
		// fmt.Println(dataPkg.Token)
	}
	// fmt.Println(r.URL.Path)
	// fmt.Println(r.URL.RawQuery)
	// fmt.Println(r.URL.Path)
	// fmt.Println(r.URL.Query()) // url values
	// fmt.Println(reflect.TypeOf(r.URL.Query()))

	// token=test``
	// /users/3/data
	// map[token:[test]]
	// url.Values
}
