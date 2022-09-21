package model

import "encoding/json"

type Pv struct {
	Id         int     `json:"Id"`
	GenkW      float64 `json:"HenkW"`
	Hz         float64 `json:"Hz"`
	Temp       float64 `json:"Temp"`
	ModuleTemp float64 `json:"ModuleTemp"`
	Time       string  `json:"Time"`
}

func (p *Pv) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&p, "", " ")
}

func (p *Pv) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, p)
}
