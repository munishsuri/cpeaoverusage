package model

type ServiceMetrics struct {
	Metric        string
	PlanId        string
	SubaccountId  string `json:",omitempty"`
	ThreholdValue int
}

type MetricsThrehsoldValue struct {
	Value          int
	PassedThrehold bool
}
