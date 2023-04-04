package model

type AlertNotificationConfig struct {
	Url           string `json:"Url"`
	Client_Id     string `json:"Client_Id"`
	Client_Secret string `json:"Client_Secret"`
}

type AlertNotificationBody struct {
	EventType string   `json:"eventType"`
	Severity  string   `json:"severity"`
	Category  string   `json:"category"`
	Body      string   `json:"body"`
	Subject   string   `json:"subject"`
	Resource  Resource `json:"resource"`
}

type Resource struct {
	ResourceName string `json:"resourceName"`
	ResourceType string `json:"resourceType"`
	Tags         Tag    `json:"tags"`
}

type Tag struct {
	Env string `json:"env"`
}
