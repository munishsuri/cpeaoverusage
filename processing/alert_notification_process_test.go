package processing

import (
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	testutils "cpea_monthly_usage/test_utils"
	"encoding/json"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_triggerAlert_Pass(t *testing.T) {

	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitAlertNotifictaion()

	// saving original env
	envOriginal := os.Getenv(env.Dest_ENV_NAME)

	os.Setenv(env.Dest_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"DestName\":\"dest\",\"Token_Url\":\"url\"}")

	// trigger exceeded alert
	triggerAlert("exceeded")

	//asserting the log
	assert.Equal(t, faekPrintClient.Args[0], "Alert Notification Response")

	bodyAlert := model.AlertNotificationBody{}
	json.Unmarshal([]byte(faekPrintClient.Args[1]), &bodyAlert)

	assert.Equal(t, bodyAlert.Body, "exceeded")

	//setting the env back
	os.Setenv(env.Dest_ENV_NAME, envOriginal)
	FMTPrintln = fmtOriginal

}

func Test_triggerAlert_FailedCall(t *testing.T) {

	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitAlertNotifictaion()

	// saving original env
	envOriginal := os.Getenv(env.Dest_ENV_NAME)

	os.Setenv(env.Dest_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"DestName\":\"dest\",\"Token_Url\":\"url\"}")

	// trigger exceeded alert
	triggerAlert("exceededFail")

	//error would have been triggered
	assert.Equal(t, faekPrintClient.Args[0], "Error While Fetching Alert Notification Response")

	//setting the env back
	os.Setenv(env.Dest_ENV_NAME, envOriginal)
	FMTPrintln = fmtOriginal

}

func Test_checkThresholdAndtriggerAlertWithSubaccount_NoThresholdExceeded(t *testing.T) {
	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln

	threshold := map[[4]string]model.MetricsThrehsoldValue{}

	threshold[[4]string{"a", "b", "c", "d"}] = model.MetricsThrehsoldValue{
		Value:          1,
		PassedThrehold: false,
	}

	checkThresholdAndtriggerAlertWithSubaccount(threshold)
	assert.Equal(t, faekPrintClient.Args[0], "Threshold Not Exceeded")

	FMTPrintln = fmtOriginal
}

func Test_checkThresholdAndtriggerAlertWithSubaccount_ThresholdExceeded(t *testing.T) {
	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitAlertNotifictaion()

	// saving original env
	envOriginal := os.Getenv(env.Dest_ENV_NAME)

	os.Setenv(env.Dest_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"DestName\":\"dest\",\"Token_Url\":\"url\"}")

	threshold := map[[4]string]model.MetricsThrehsoldValue{}

	threshold[[4]string{"a", "b", "c", "d"}] = model.MetricsThrehsoldValue{
		Value:          1,
		PassedThrehold: true,
	}

	checkThresholdAndtriggerAlertWithSubaccount(threshold)
	assert.Equal(t, faekPrintClient.Args[0], "Alert Notification Response")

	bodyAlert := model.AlertNotificationBody{}
	json.Unmarshal([]byte(faekPrintClient.Args[1]), &bodyAlert)

	assert.Equal(t, bodyAlert.Body, " Service Exceed - a, Metric Exceeded - b, Plan Exceeded - c, Sub account - d, Current Consumption - 1------")
	// setting the env back
	os.Setenv(env.Dest_ENV_NAME, envOriginal)

	FMTPrintln = fmtOriginal
}

func Test_checkThresholdAndtriggerAlert_NoThresholdExceeded(t *testing.T) {
	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln

	threshold := map[[3]string]model.MetricsThrehsoldValue{}

	threshold[[3]string{"a", "b", "c"}] = model.MetricsThrehsoldValue{
		Value:          1,
		PassedThrehold: false,
	}

	checkThresholdAndtriggerAlert(threshold)
	assert.Equal(t, faekPrintClient.Args[0], "Threshold Not Exceeded")

	FMTPrintln = fmtOriginal
}

func Test_checkThresholdAndtriggerAlert_ThresholdExceeded(t *testing.T) {
	fmtOriginal := FMTPrintln
	faekPrintClient := testutils.FakePrint{}
	FMTPrintln = faekPrintClient.FakePrintln
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitAlertNotifictaion()

	// saving original env
	envOriginal := os.Getenv(env.Dest_ENV_NAME)

	os.Setenv(env.Dest_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"DestName\":\"dest\",\"Token_Url\":\"url\"}")

	threshold := map[[3]string]model.MetricsThrehsoldValue{}

	threshold[[3]string{"a", "b", "c"}] = model.MetricsThrehsoldValue{
		Value:          1,
		PassedThrehold: true,
	}

	checkThresholdAndtriggerAlert(threshold)
	assert.Equal(t, faekPrintClient.Args[0], "Alert Notification Response")

	bodyAlert := model.AlertNotificationBody{}
	json.Unmarshal([]byte(faekPrintClient.Args[1]), &bodyAlert)

	assert.Equal(t, bodyAlert.Body, " Service Exceed - a, Metric Exceeded - b, Plan Exceeded - c, Current Consumption - 1------")
	// setting the env back
	os.Setenv(env.Dest_ENV_NAME, envOriginal)

	FMTPrintln = fmtOriginal
}
