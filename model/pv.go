package model

import "encoding/json"

type Pv struct {
	Id         int     `json:"id"`
	GenkW      float64 `json:"genkW"`
	Hz         float64 `json:"hz"`
	Temp       float64 `json:"temp"`
	ModuleTemp float64 `json:"moduleTemp"`
	Time       string  `json:"time"`
}

func (p *Pv) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&p, "", " ")
}

func (p *Pv) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, p)
}
