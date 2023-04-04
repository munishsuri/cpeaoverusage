package testutils

import (
	"cpea_monthly_usage/model"
	"encoding/json"
)

func GetUasRespTest() model.UASMonthlyUsageResponse {

	uasRespString := `{
		"content": [
			{
				"subaccountId": "s1",
				"serviceId": "html-repo",
				"plan": "app-host",
				"measureId": "mb_storage",
				"usage":5
			},
			{
				"subaccountId": "s1",
				"serviceId": "html-repo",
				"plan": "app-host",
				"measureId": "mb_storage",
				"usage": 5
			},
			{
				"subaccountId": "s2",
				"serviceId": "html-repo",
				"plan": "app-host",
				"measureId": "mb_storage",
				"usage": 5
			}
			
		]
	}`

	var uasResp model.UASMonthlyUsageResponse

	json.Unmarshal([]byte(uasRespString), &uasResp)

	return uasResp
}
