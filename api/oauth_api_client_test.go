package api

import (
	testutils "cpea_monthly_usage/test_utils"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_Oauth_Post_Pass(t *testing.T) {

	testutils.InitOAuth()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "urlOauth"

	oauthClient := GetAPIClient(url, &http.Client{})
	body, err := oauthClient.PostAPI([]byte("verifybody"), "/testPOSTPass")

	if err != nil {
		t.Error("Error occured")
	}

	if string(body) != string("verifybody") {
		t.Error("Error occured")
	}

}

func Test_Oauth_Post_Fail(t *testing.T) {

	testutils.InitOAuth()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "urlOauth"

	oauthClient := GetAPIClient(url, &http.Client{})
	_, err := oauthClient.PostAPI([]byte("verifybody"), "/testPOSTFail")

	if err == nil {
		t.Error("Error occured")
	}

}

func Test_Oauth_GET_Pass(t *testing.T) {

	testutils.InitOAuth()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "urlOauth"

	oauthClient := GetAPIClient(url, &http.Client{})
	_, err := oauthClient.GetAPI("/testGETPass")

	if err != nil {
		t.Error("Error occured")
	}
}

func Test_Oauth_Get_Fail(t *testing.T) {

	testutils.InitOAuth()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "urlOauth"

	oauthClient := GetAPIClient(url, &http.Client{})
	_, err := oauthClient.GetAPI("/testGETFail")

	if err == nil {
		t.Error("Error occured")
	}

}
