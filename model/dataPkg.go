package model

import "encoding/json"

type DataPkg struct {
	UserId   int    `json:"UserId"`
	Token    string `json:"Token"`
	JsonData []byte `json:"JsonData"`
}

func (d *DataPkg) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&d, "", " ")
}

func (d *DataPkg) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, &d)
}
