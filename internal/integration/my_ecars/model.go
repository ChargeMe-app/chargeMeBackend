package my_ecars

type MyECarsConnector struct {
	Current interface{} `json:"current"`
	Phase   int         `json:"phase"`
	Voltage interface{} `json:"voltage"`
	Power   interface{} `json:"power"`
	Type    interface{} `json:"type"`
	State   interface{} `json:"state"`
	Cost    interface{} `json:"coast"`
}

type MyECarsStation struct {
	Id         string             `json:"id"`
	Online     int                `json:"online"`
	Name       string             `json:"name"`
	Address    string             `json:"address"`
	Location   string             `json:"location"`
	Reserv     int                `json:"reserv"`
	Phone      string             `json:"phone"`
	Access     string             `json:"access"`
	Sleep      interface{}        `json:"sleep"`
	Connectors []MyECarsConnector `yaml:"connectors"`
}

type MyECarsStationsResponse struct {
	Status string           `json:"status"`
	Error  string           `json:"error"`
	Evse   []MyECarsStation `json:"evse"`
}
