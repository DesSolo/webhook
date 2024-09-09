package entities

type Request struct {
	UUID    string              `json:"uuid"`
	Token   string              `json:"token"`
	Date    string              `json:"date"`
	IP      string              `json:"ip"`
	Method  string              `json:"method"`
	URI     string              `json:"uri"`
	Query   string              `json:"query"`
	Headers map[string][]string `json:"headers"`
	Data    []byte              `json:"data"`
}
