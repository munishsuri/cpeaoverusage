package model

type DestConfig struct {
	Token_Url     string `json:"Token_Url"`
	Client_Id     string `json:"Client_Id"`
	Client_Secret string `json:"Client_Secret"`
	URL           string `json:"Url"`
	DestName      string `json:"DestName"`
}

type Destination struct {
	Owner struct {
		SubaccountID string      `json:"SubaccountId"`
		InstanceID   interface{} `json:"InstanceId"`
	} `json:"owner"`
	DestinationConfiguration DestinationConfiguration `json:"destinationConfiguration"`
	AuthTokens               []DestHttpHeader         `json:"authTokens"`
}

type DestinationConfiguration struct {
	Name           string `json:"Name"`
	Type           string `json:"Type"`
	URL            string `json:"URL"`
	Authentication string `json:"Authentication"`
	ProxyType      string `json:"ProxyType"`
	User           string `json:"User"`
	Password       string `json:"Password"`
}

type DestHttpHeader struct {
	Type       string `json:"type"`
	Value      string `json:"value"`
	HTTPHeader struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"http_header"`
}
