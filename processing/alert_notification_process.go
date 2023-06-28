package processing

import (
	"bytes"
	"context"
	"cpea_monthly_usage/api"
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/oauth2/clientcredentials"
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
	destConfig, err := env.GetDestConfigValue()

	if err != nil {
		fmt.Println("Error In getting Alert Notification Config")
		env.EndChan <- 1
	}

	// destination client
	destClientConfig := clientcredentials.Config{
		ClientID:     destConfig.Client_Id,
		ClientSecret: destConfig.Client_Secret,
		TokenURL:     destConfig.Token_Url + "/oauth/token",
	}
	alertClient := api.GetAPIClient(destConfig.URL, destClientConfig.Client(context.Background()))

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

	destBytes, e := alertClient.GetAPI(env.DestConfigURL + destConfig.DestName)

	if e != nil {
		FMTPrintln("Error While Fetching Destination", e)
		return
	}

	// send event via request
	var destResponse model.Destination
	err = json.Unmarshal(destBytes, &destResponse)

	r, _ := http.NewRequest("POST", destResponse.DestinationConfiguration.URL, bytes.NewReader(payloadJson))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set(destResponse.AuthTokens[0].HTTPHeader.Key, destResponse.AuthTokens[0].HTTPHeader.Value)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		FMTPrintln("Error While Fetching Alert Notification Response", err)
		return
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)

	FMTPrintln("Alert Notification Response", string(b))
}
