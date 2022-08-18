package model

import "encoding/json"

type Pv struct {
	Id         int
	GenkW      float64
	Hz         float64
	Temp       float64
	ModuleTemp float64
}

func (p *Pv) MarshalJson() ([]byte, error) {
	return json.Marshal(&p)
}
