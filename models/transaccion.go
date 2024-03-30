package models

type Comercio struct {
	Monto  float64 `json:"monto"`
	Rif    string  `json:"rif"`
	Banco  string  `json:"banco"`
	Phone  string  `json:"phone"`
}

type Usuario struct {
	Banco  string `json:"banco"`
	Cedula string `json:"cedula"`
	Phone  string `json:"phone"`
}

type TransaccionInput struct {
	Comercio Comercio `json:"Comercio"`
	Usuario  Usuario  `json:"Usuario"`
}
