package model

import "encoding/json"

// not used
type DataPkg struct {
	UserId int `json:"UserId"`
	// UserNm string `json:"UserNm"`
	// Token    string `json:"Token"`
	JsonData Pv `json:"JsonData"`
}

func (d *DataPkg) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&d, "", " ")
}

func (d *DataPkg) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, &d)
}
