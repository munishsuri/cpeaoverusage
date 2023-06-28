package testutils

import (
	"cpea_monthly_usage/model"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/jarcoal/httpmock"
)

func InitBasic() {
	httpmock.RegisterResponder("POST", "url/testPassBasic",
		func(req *http.Request) (*http.Response, error) {
			b, _ := ioutil.ReadAll(req.Body)
			return httpmock.NewBytesResponse(200, b), nil
		})

	httpmock.RegisterResponder("POST", "url/testFailBasic",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Error Test")
		})

}

func InitOAuth() {
	httpmock.RegisterResponder("GET", "urlOauth/testGETPass",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, "Passed"), nil
		})

	httpmock.RegisterResponder("GET", "urlOauth/testGETFail",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Error Test")
		})
	httpmock.RegisterResponder("POST", "urlOauth/testPOSTPass",
		func(req *http.Request) (*http.Response, error) {
			b, _ := ioutil.ReadAll(req.Body)
			return httpmock.NewBytesResponse(200, b), nil
		})

	httpmock.RegisterResponder("POST", "urlOauth/testPOSTFail",
		func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Error Test")
		})
}

func InitAlertNotifictaion() {
	httpmock.RegisterResponder("POST", "url/oauth/token",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{"access_token":"token"}`), nil

		})
	httpmock.RegisterResponder("GET", "url/destination-configuration/v1/destinations/dest",
		func(req *http.Request) (*http.Response, error) {
			body, _ := json.Marshal(model.Destination{
				DestinationConfiguration: model.DestinationConfiguration{
					URL: "testurldest",
				},
				AuthTokens: []model.DestHttpHeader{
					model.DestHttpHeader{
						HTTPHeader: struct {
							Key   string "json:\"key\""
							Value string "json:\"value\""
						}{},
					},
				},
			})
			return httpmock.NewBytesResponse(200, body), nil

		})
	httpmock.RegisterResponder("POST", "testurldest",
		func(req *http.Request) (*http.Response, error) {
			b, _ := ioutil.ReadAll(req.Body)
			var alertBody model.AlertNotificationBody
			json.Unmarshal(b, &alertBody)
			if alertBody.Body == "exceededFail" {
				return httpmock.NewStringResponse(400, string(b)), errors.New("Fail")
			}
			return httpmock.NewStringResponse(200, string(b)), nil

		})
}

func InitUAS() {

	httpmock.RegisterResponder("GET", "=~^url/reports/v1/*",
		func(req *http.Request) (*http.Response, error) {

			return httpmock.NewStringResponse(200, "{}"), nil

		})

	httpmock.RegisterResponder("POST", "url/oauth/token",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{"access_token":"token"}`), nil

		})

	httpmock.RegisterResponder("GET", "=~^urlBadBody/reports/v1/*",
		func(req *http.Request) (*http.Response, error) {

			return httpmock.NewStringResponse(200, ""), nil

		})
	httpmock.RegisterResponder("GET", "=~^urlBadUas/reports/v1/*",
		func(req *http.Request) (*http.Response, error) {

			return httpmock.NewStringResponse(200, ""), errors.New("error in uas")

		})
}

func InitThreshold() {

	httpmock.RegisterResponder("GET", "=~^url/reports/v1/*",
		func(req *http.Request) (*http.Response, error) {
			a, _ := json.Marshal(GetUasRespTest())
			return httpmock.NewStringResponse(200, string(a)), nil

		})

	httpmock.RegisterResponder("POST", "url/oauth/token",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{"access_token":"token"}`), nil

		})

}
