package processing

import (
	"context"
	"cpea_monthly_usage/api"
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/oauth2/clientcredentials"
)

var Chan = env.EndChan
var Println = fmt.Println

func getUASResponse() model.UASMonthlyUsageResponse {
	uasConfigValues, err := env.GetUASConfigValue()

	// getting uas config
	if err != nil {
		fmt.Println(err)
		env.EndChan <- 1

	}

	//setting uas config
	uasClientOauthConfig := clientcredentials.Config{
		ClientID:     uasConfigValues.Client_Id,
		ClientSecret: uasConfigValues.Client_Secret,
		TokenURL:     uasConfigValues.Token_Url + "/oauth/token",
	}

	// uas client get
	uasClient := api.GetAPIClient(uasConfigValues.URL, uasClientOauthConfig.Client(context.Background()))

	// executing call for uas current month
	url := "/reports/v1/monthlyUsage?fromDate="

	year, month, _ := time.Now().Date()
	// date format for API
	dateAPI := year*100 + int(month)
	url = url + strconv.Itoa(dateAPI) + "&toDate=" + strconv.Itoa(dateAPI)

	responseUASBytes, err := uasClient.GetAPI(url)
	// parse the response into UAS Variable
	var uasResponse model.UASMonthlyUsageResponse

	if err != nil {
		Println("Error Occurred While Getting UAS Response", err)
		Chan <- 1
		return uasResponse
	}

	err = json.Unmarshal(responseUASBytes, &uasResponse)

	if err != nil {
		Println("Error Occurred While Parsing Response", err)
		Chan <- 2
	}

	return uasResponse

}
