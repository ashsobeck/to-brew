package types

type Bean struct {
	Id string `json:"id"`
	Name     string `json:"name"`
	Roaster  string `json:"roaster"`
	Country  string `json:"country"`
	Varietal string `json:"varietal"`
	Process  string `json:"process"`
	Altitude string `json:"altitude"`
	Notes    string `json:"notes"`
	Weight   float32 `json:"weight"`
}
