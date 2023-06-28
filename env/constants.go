package env

import "cpea_monthly_usage/model"

var EndChan = make(chan int)

const (
	SERVICES_ENV_NAME        = "Services"
	TIME_ENV_NAME            = "Time"
	UAS_ENV_NAME             = "Uas"
	Event_TYPE_ENV_NAME      = "Event"
	MODE_SUBACCOUNT_ENV_NAME = "SubaccountMode"
	Dest_ENV_NAME            = "Destination"
)

// Service_ID , MEASURE_ID
var Services_Value_Map_Constants = map[string][]model.ServiceMetrics{
	"CloudIntegration": []model.ServiceMetrics{model.ServiceMetrics{
		Metric:        "connections",
		ThreholdValue: 100,
		PlanId:        "standard",
	},
		model.ServiceMetrics{
			Metric:        "tenants",
			ThreholdValue: 5,
			PlanId:        "standard",
		}},
	"data-intelligence": []model.ServiceMetrics{model.ServiceMetrics{
		Metric:        "capacity_units",
		ThreholdValue: 100,
		PlanId:        "enterprise",
	}},
	"sap-workzone": []model.ServiceMetrics{model.ServiceMetrics{
		Metric:        "swz_users",
		ThreholdValue: 5,
		PlanId:        "standard",
	}, model.ServiceMetrics{
		Metric:        "swz_connections",
		ThreholdValue: 5,
		PlanId:        "standard",
	}},
}

const (
	Time_Constant  = 5
	EventName      = "cpeacreditsover"
	DestConfigURL  = "/destination-configuration/v1/destinations/"
	SubaccountMode = false
)
