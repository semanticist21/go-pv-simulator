package model

import "encoding/json"

type DataPkg struct {
	UserId   int    `json:"userId"`
	Token    string `json:"token"`
	JsonData []byte `json:"jsonData"`
}

func (d *DataPkg) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&d, "", " ")
}

func (d *DataPkg) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, &d)
}
