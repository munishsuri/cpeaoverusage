package processing

import (
	"cpea_monthly_usage/env"
	testutils "cpea_monthly_usage/test_utils"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_process_subaccount(t *testing.T) {
	timeValOrg := timeVal
	processFuncOrg := processFunc

	time.Sleep(2 * time.Second)
	timeVal = time.Duration(2) * time.Second
	// creating fake process
	fakeProcess := testutils.FakeProcess{}
	processFunc = fakeProcess.ProcessFunTest

	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "false")

	// executing test
	CheckMonthlyData()
	time.Sleep(3 * time.Second)
	assert.Equal(t, fakeProcess.Called, 1)

	timeVal = timeValOrg
	processFunc = processFuncOrg
}

func Test_process_withsubaccount(t *testing.T) {
	timeValOrg := timeVal
	processFuncOrg := processFuncSubaccount

	time.Sleep(2 * time.Second)

	os.Setenv(env.MODE_SUBACCOUNT_ENV_NAME, "true")
	timeVal = time.Duration(2) * time.Second

	// creating fake process
	fakeProcess := testutils.FakeProcess{}
	processFuncSubaccount = fakeProcess.ProcessFunTest

	// executing test
	CheckMonthlyData()
	time.Sleep(3 * time.Second)
	assert.Equal(t, fakeProcess.Called, 1)

	timeVal = timeValOrg
	processFuncSubaccount = processFuncOrg
}

func Test_CheckErrorAndEndApplication(t *testing.T) {
	endChannelorg := endChannel
	endChannel = make(chan int)

	fakeExit := testutils.FakeExit{}
	osexit = fakeExit.Exit

	go CheckErrorAndEndApplication()
	endChannel <- 1

	time.Sleep(time.Second)
	assert.Equal(t, fakeExit.Code, 1)

	endChannel = endChannelorg

}
