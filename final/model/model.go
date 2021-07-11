package model

type Alienigena struct {
	Id       int    `json:"id"`
	Longitud string `json:"longitud"`
	Latitud  string `json:"latitud"`
	Nombre   string `json:"nombre"`
}

type Aliado struct {
	Id       int    `json:"id"`
	Latitud  string `json:"latitud"`
	Longitud string `json:"longitud"`
	Nombre   string `json:"nombre"`
}
