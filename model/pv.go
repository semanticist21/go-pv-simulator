package model

import "encoding/json"

type Pv struct {
	PvId       int     `json:"PvId"`
	Time       string  `json:"Time"`
	GenkW      float64 `json:"GenkW"`
	Hz         float64 `json:"Hz"`
	Temp       float64 `json:"Temp"`
	ModuleTemp float64 `json:"ModuleTemp"`
}

func (p *Pv) MarshalJson() ([]byte, error) {
	return json.MarshalIndent(&p, "", " ")
}

func (p *Pv) UnMarshalJson(data []byte) error {
	return json.Unmarshal(data, p)
}
