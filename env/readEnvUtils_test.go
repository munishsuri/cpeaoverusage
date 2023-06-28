package env

import (
	"cpea_monthly_usage/model"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetService_Blank_os(t *testing.T) {

	envOrg := os.Getenv(SERVICES_ENV_NAME)
	os.Setenv(SERVICES_ENV_NAME, "")

	envval, err := GetService()

	assert.Equal(t, err, nil)
	assert.Equal(t, envval, Services_Value_Map_Constants)

	os.Setenv(SERVICES_ENV_NAME, envOrg)

}

func Test_GetService_Error(t *testing.T) {

	envOrg := os.Getenv(SERVICES_ENV_NAME)
	os.Setenv(SERVICES_ENV_NAME, "fail")

	envval, err := GetService()
	fmt.Println(envval, err)
	assert.NotEqual(t, err, nil)

	os.Setenv(SERVICES_ENV_NAME, envOrg)

}

func Test_GetService_EnvVal(t *testing.T) {

	envOrg := os.Getenv(SERVICES_ENV_NAME)
	os.Setenv(SERVICES_ENV_NAME, `{
		"di": [
		  {
			"Metric": "cu",
			"ThreholdValue": 1000,
			"PlanId": "enterprise"
		  }
		]
	}`)

	envval, err := GetService()

	assert.Equal(t, err, nil)
	assert.Equal(t, envval, map[string][]model.ServiceMetrics{
		"di": []model.ServiceMetrics{{
			Metric:        "cu",
			ThreholdValue: 1000,
			PlanId:        "enterprise",
		}},
	})

	os.Setenv(SERVICES_ENV_NAME, envOrg)
}

func Test_GetUASConfigValue_error(t *testing.T) {

	envOrg := os.Getenv(UAS_ENV_NAME)
	os.Setenv(UAS_ENV_NAME, "")

	envval, err := GetUASConfigValue()

	assert.NotEqual(t, err, nil)
	assert.Equal(t, model.UASConfig{}, envval)

	os.Setenv(UAS_ENV_NAME, envOrg)

}

func Test_GetUASConfigValue_error_Json_Parse(t *testing.T) {

	envOrg := os.Getenv(UAS_ENV_NAME)
	os.Setenv(UAS_ENV_NAME, "fail")

	envval, err := GetUASConfigValue()

	assert.NotEqual(t, err, nil)
	assert.Equal(t, model.UASConfig{}, envval)

	os.Setenv(UAS_ENV_NAME, envOrg)

}

func Test_GetUASConfigValue_Pass(t *testing.T) {

	envOrg := os.Getenv(UAS_ENV_NAME)
	os.Setenv(UAS_ENV_NAME, `
	{"Client_Id":"id",
	"Client_Secret":"secret",
	"Token_Url":"token_url",
	"Url":"url"}
	`)

	envval, err := GetUASConfigValue()

	assert.Equal(t, err, nil)
	assert.Equal(t, model.UASConfig{
		Token_Url:     "token_url",
		Client_Id:     "id",
		Client_Secret: "secret",
		URL:           "url",
	}, envval)

	os.Setenv(UAS_ENV_NAME, envOrg)

}

func Test_GetTimeInMinutes_Pass(t *testing.T) {
	envOrg := os.Getenv(TIME_ENV_NAME)

	os.Setenv(TIME_ENV_NAME, "10")

	timeCheck := GetTimeInMinutes()

	assert.Equal(t, 10*time.Minute, timeCheck)

	os.Setenv(TIME_ENV_NAME, envOrg)

}

func Test_GetTimeInMinutes_ErrorString(t *testing.T) {
	envOrg := os.Getenv(TIME_ENV_NAME)

	os.Setenv(TIME_ENV_NAME, "{}")

	timeCheck := GetTimeInMinutes()

	assert.Equal(t, 5*time.Minute, timeCheck)

	os.Setenv(TIME_ENV_NAME, envOrg)

}

func Test_GetTimeInMinutes(t *testing.T) {
	envOrg := os.Getenv(TIME_ENV_NAME)

	os.Setenv(TIME_ENV_NAME, "")

	timeCheck := GetTimeInMinutes()

	assert.Equal(t, 5*time.Minute, timeCheck)

	os.Setenv(TIME_ENV_NAME, envOrg)

}

func Test_getEvent_name(t *testing.T) {

	envOrg := os.Getenv(Event_TYPE_ENV_NAME)

	os.Setenv(Event_TYPE_ENV_NAME, "")

	event := GetEventName()

	assert.Equal(t, event, "cpeacreditsover")

	os.Setenv(Event_TYPE_ENV_NAME, envOrg)

}

func Test_getEvent_name_custom(t *testing.T) {

	envOrg := os.Getenv(Event_TYPE_ENV_NAME)

	os.Setenv(Event_TYPE_ENV_NAME, "custom")

	event := GetEventName()

	assert.Equal(t, event, "custom")

	os.Setenv(Event_TYPE_ENV_NAME, envOrg)

}

func Test_GetSubaccountMode_false(t *testing.T) {

	envOrg := os.Getenv(MODE_SUBACCOUNT_ENV_NAME)

	os.Setenv(MODE_SUBACCOUNT_ENV_NAME, "")

	event := GetSubaccountMode()

	assert.Equal(t, event, false)

	os.Setenv(MODE_SUBACCOUNT_ENV_NAME, envOrg)

}

func Test_GetSubaccountMode_true(t *testing.T) {

	envOrg := os.Getenv(MODE_SUBACCOUNT_ENV_NAME)

	os.Setenv(MODE_SUBACCOUNT_ENV_NAME, "true")

	event := GetSubaccountMode()

	assert.Equal(t, event, true)

	os.Setenv(MODE_SUBACCOUNT_ENV_NAME, envOrg)

}

func Test_GetDestConfigValue_true(t *testing.T) {

	envOrg := os.Getenv(Dest_ENV_NAME)

	os.Setenv(Dest_ENV_NAME, `
	{
		"Client_Id":"id",
		"Client_Secret":"secret",
		"Token_Url":"token",
		"Url":"url",
		"DestName":"destination"
	}`)

	destConfig, _ := GetDestConfigValue()

	assert.Equal(t, destConfig, model.DestConfig{
		Token_Url:     "token",
		Client_Id:     "id",
		Client_Secret: "secret",
		URL:           "url",
		DestName:      "destination",
	})

	os.Setenv(Dest_ENV_NAME, envOrg)

}
