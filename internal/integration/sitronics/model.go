package sitronics

type SitronicsConnector struct {
	ConnectorId                int    `json:"ConnectorId"`
	Type                       string `json:"Type"`
	TypeName                   string `json:"TypeName"`
	StatusName                 string `json:"StatusName"`
	MaxPower                   string `json:"MaxP"`
	Status                     string `json:"Status"`
	StatusInt                  int    `json:"StatusInt"`
	TimeToEndChargingSpecified bool   `json:"TimeToEndChargingSpecified"`
}

type SitronicsService struct {
	Name           string `json:"Name"`
	Unit           string `json:"Unit"`
	Price          int    `json:"Price"`
	CurrencySymbol string `json:"CurrencySymbol"`
	CurrencyName   string `json:"CurrencyName"`
	BankId         int    `json:"BankId"`
}

type SitronicsAdvancedProperty struct {
	Cafe         bool `json:"cafe"`
	CloseAccess  bool `json:"closeAccess"`
	Hotel        bool `json:"hotel"`
	Park         bool `json:"park"`
	Parking      bool `json:"parking"`
	Shop         bool `json:"shop"`
	ShoppingMall bool `json:"shoppingMall"`
	Wc           bool `json:"wc"`
	Wifi         bool `json:"wifi"`
	Hour24       bool `json:"hour24"`
}

type SitronicsStation struct {
	Name                  string                    `json:"Name"`
	Id                    string                    `json:"Id"`
	NodeId                string                    `json:"NodeId"`
	IdInt                 int                       `json:"IdInt"`
	Type                  string                    `json:"Type"`
	Address               string                    `json:"Address"`
	IpAddress             string                    `json:"IpAddress"`
	Status                int                       `json:"Status"`
	Latitude              float32                   `json:"Latitude"`
	Longitude             float32                   `json:"Longitude"`
	City                  string                    `json:"City"`
	PhotoUrl              *string                   `json:"PhotoUrl"`
	Owner                 string                    `json:"Owner"`
	PhoneNumber           *string                   `json:"PhoneNumber"`
	Manufacturer          *string                   `json:"Manufacturer"`
	Connectors            []SitronicsConnector      `json:"Connectors"`
	ServiceList           []SitronicsService        `json:"ServiceList"`
	AdvancedProperty      SitronicsAdvancedProperty `json:"AdvancedProperty"`
	PublicDescription     string                    `json:"PublicDescription"`
	Commentary            string                    `json:"Commentary"`
	WarningText           string                    `json:"WarningText"`
	PassUrl               string                    `json:"PassUrl"`
	WorkingTime           string                    `json:"WorkingTime"`
	MulticonnectorSupport bool                      `json:"MulticonnectorSupport"`
	TimeLimit             bool                      `json:"TimeLimit"`
	MaxPower              int                       `json:"MaxPower"`
}

type SitronicsMapInfo struct {
	CPList []SitronicsStation `yaml:"CPList"`
}

type SitronicsStationInfo struct {
	CPCard SitronicsStation `yaml:"CPCard"`
}
