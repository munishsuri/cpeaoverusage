package processing

import (
	"cpea_monthly_usage/env"
	testutils "cpea_monthly_usage/test_utils"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_getThresholdResponse_true(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitThreshold()
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)

	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")

	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 10,
			"PlanId": "app-host"
		  }
		]	
	  }
	`)

	// executing test
	threshold := getThresholdResponse()

	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].PassedThrehold, true)
	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].Value, 15)

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_getThresholdResponse_fail(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitThreshold()
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)

	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")

	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 20,
			"PlanId": "app-host"
		  }
		]	
	  }
	`)

	// executing test
	threshold := getThresholdResponse()

	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].PassedThrehold, false)

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_getThresholdResponseWithsubaccount_fail(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitThreshold()
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)

	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")

	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 20,
			"PlanId": "app-host",
			"SubaccountId":"s1"
		  }
		]	
	  }
	`)

	// executing test
	threshold := getThresholdResponseWithSubaccount()

	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s1"}].PassedThrehold, false)

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_getThresholdResponseWithsubaccount_pass(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitThreshold()
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)

	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")

	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 5,
			"PlanId": "app-host",
			"SubaccountId":"s1"
		  }
		]	
	  }
	`)

	// executing test
	threshold := getThresholdResponseWithSubaccount()

	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s1"}].PassedThrehold, true)
	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s1"}].Value, 10)

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}
func Test_checkThreshold_true(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan
	//
	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 10,
			"PlanId": "app-host"
		  }
		]	
	  }
	`)
	threshold := checkThreshold(testutils.GetUasRespTest())

	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].PassedThrehold, true)
	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].Value, 15)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_checkThreshold_false(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan
	//
	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 20,
			"PlanId": "app-host"
		  }
		]	
	  }
	`)
	threshold := checkThreshold(testutils.GetUasRespTest())
	assert.Equal(t, threshold[[3]string{"html-repo", "mb_storage", "app-host"}].PassedThrehold, false)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_checkThreshold_withsubaccount_true(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan
	//
	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "true")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 5,
			"PlanId": "app-host",
			"SubaccountId":"s1"
		  }
		]	
	  }
	`)
	threshold := checkThresholdWithSubaccount(testutils.GetUasRespTest())

	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s1"}].PassedThrehold, true)
	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s1"}].Value, 10)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}

func Test_checkThreshold_withsubaccount_false(t *testing.T) {
	eChan := make(chan int)
	go checkChan(eChan)
	EndChan = eChan
	//
	ser := os.Getenv(env.SERVICES_ENV_NAME)
	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "true")
	os.Setenv(env.SERVICES_ENV_NAME, `{
		"html-repo": [
		  {
			"Metric": "mb_storage",
			"ThreholdValue": 5,
			"PlanId": "app-host",
			"SubaccountId":"s2"
		  }
		]	
	  }
	`)
	threshold := checkThresholdWithSubaccount(testutils.GetUasRespTest())

	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s2"}].PassedThrehold, false)
	assert.Equal(t, threshold[[4]string{"html-repo", "mb_storage", "app-host", "s2"}].Value, 5)
	EndChan = env.EndChan
	os.Setenv(env.SERVICES_ENV_NAME, ser)

}
