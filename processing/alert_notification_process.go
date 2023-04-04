package processing

import (
	"cpea_monthly_usage/api"
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var FMTPrintln = fmt.Println

func checkThresholdAndtriggerAlert(thresholds map[[3]string]model.MetricsThrehsoldValue) {

	overusage := false
	exceededValues := ""

	// checking if threshold is exceeding
	for k, v := range thresholds {

		if v.PassedThrehold {
			overusage = true
			exceededValues = exceededValues + " Service Exceed - " + k[0] + ", Metric Exceeded - " + k[1] + ", Plan Exceeded - " + k[2] + ", Current Consumption - " + strconv.Itoa(v.Value) + "------"
		}

	}

	if !overusage {
		FMTPrintln("Threshold Not Exceeded")
		return
	}

	triggerAlert(exceededValues)

}

func checkThresholdAndtriggerAlertWithSubaccount(thresholds map[[4]string]model.MetricsThrehsoldValue) {

	overusage := false
	exceededValues := ""

	// checking if threshold is exceeding
	for k, v := range thresholds {

		if v.PassedThrehold {
			overusage = true
			exceededValues = exceededValues + " Service Exceed - " + k[0] + ", Metric Exceeded - " + k[1] + ", Plan Exceeded - " + k[2] + ", Sub account - " + k[3] + ", Current Consumption - " + strconv.Itoa(v.Value) + "------"
		}

	}

	if !overusage {
		FMTPrintln("Threshold Not Exceeded")
		return
	}

	triggerAlert(exceededValues)

}

func triggerAlert(exceededValues string) {
	// Trigger Alert
	alertConfig, err := env.GetAlertConfig()
	if err != nil {
		fmt.Println("Error In getting Alert Notification Config")
		env.EndChan <- 1
	}

	alertClient := api.GetBasicAPIClient(alertConfig.Url, alertConfig.Client_Id, alertConfig.Client_Secret, &http.Client{})

	// create payload for request
	payload := model.AlertNotificationBody{
		Category:  "ALERT",
		Severity:  "FATAL",
		EventType: env.GetEventName(),
		Body:      exceededValues,
		Subject:   "Cloud Credits Over Exceeded",
		Resource: model.Resource{
			ResourceName: "alert-notification-app",
			ResourceType: "app",
			Tags: model.Tag{
				Env: "prod",
			},
		},
	}

	payloadJson, _ := json.Marshal(payload)

	r, e := alertClient.PostAPI(payloadJson, "/cf/producer/v1/resource-events")

	if e != nil {
		FMTPrintln("Error While Fetching Alert Notification Response", e)
		return
	}

	FMTPrintln("Alert Notification Response", string(r))
}
