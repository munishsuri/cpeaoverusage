package processing

import (
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	testutils "cpea_monthly_usage/test_utils"
	"fmt"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func Test_getUASResponse_Pass(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitUAS()

	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)

	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"url\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")
	response := getUASResponse()
	assert.Equal(t, response, model.UASMonthlyUsageResponse{})

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)

}

func Test_getUASResponse_Fail_Bad_response(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitUAS()

	fakePrintln := testutils.FakePrint{}
	Println = fakePrintln.FakePrintln

	testChan := make(chan int)
	Chan = testChan
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)
	go checkChan(testChan)
	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"urlBadBody\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")
	response := getUASResponse()

	assert.Equal(t, response, model.UASMonthlyUsageResponse{})
	assert.Equal(t, fakePrintln.Args[0], "Error Occurred While Parsing Response")

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	Chan = env.EndChan
	Println = fmt.Println

}

func Test_getUASResponse_Fail_Bad_response_UAS(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	testutils.InitUAS()

	fakePrintln := testutils.FakePrint{}
	Println = fakePrintln.FakePrintln

	testChan := make(chan int)
	Chan = testChan
	// saving original env
	envOriginal := os.Getenv(env.UAS_ENV_NAME)
	go checkChan(testChan)
	os.Setenv(env.UAS_ENV_NAME,
		"{\"Url\":\"urlBadUas\",\"Client_id\":\"user\",\"Client_Secret\":\"pass\",\"Token_Url\":\"url\"}")
	response := getUASResponse()

	assert.Equal(t, response, model.UASMonthlyUsageResponse{})
	assert.Equal(t, fakePrintln.Args[0], "Error Occurred While Getting UAS Response")

	// setting the env back
	os.Setenv(env.ALERT_NOTIFICATION_ENV_NAME, envOriginal)
	Chan = env.EndChan
	Println = fmt.Println

}

func checkChan(testChan chan int) {
	for {
		select {
		case <-testChan:
			return
		}
	}
}
