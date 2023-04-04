package api

import (
	testutils "cpea_monthly_usage/test_utils"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Basic_Auth_Post(t *testing.T) {

	testutils.InitBasic()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "url"
	user := "user"
	pass := "pass"

	basicClient := GetBasicAPIClient(url, user, pass, &http.Client{})
	body, err := basicClient.PostAPI([]byte("verifybody"), "/testPassBasic")

	if err != nil {
		t.Error("Error occured")
	}

	if string(body) != string("verifybody") {
		t.Error("Error occured")
	}

}

func Test_Basic_Auth_Post_Fail(t *testing.T) {

	testutils.InitBasic()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "url"
	user := "user"
	pass := "pass"

	basicClient := GetBasicAPIClient(url, user, pass, &http.Client{})
	_, err := basicClient.PostAPI([]byte("verifybody"), "/testFailBasic")

	if err == nil {
		t.Error("Error occured")
	}

}
