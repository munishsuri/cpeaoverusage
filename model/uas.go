package model

type UASConfig struct {
	Token_Url     string `json:"Token_Url"`
	Client_Id     string `json:"Client_Id"`
	Client_Secret string `json:"Client_Secret"`
	URL           string `json:"Url"`
}

type UASMonthlyUsageResponse struct {
	Content []struct {
		GlobalAccountID         string      `json:"globalAccountId"`
		GlobalAccountName       string      `json:"globalAccountName"`
		SubaccountID            string      `json:"subaccountId"`
		SubaccountName          string      `json:"subaccountName"`
		DirectoryID             interface{} `json:"directoryId"`
		DirectoryName           interface{} `json:"directoryName"`
		ReportYearMonth         int         `json:"reportYearMonth"`
		ServiceID               string      `json:"serviceId"`
		ServiceName             string      `json:"serviceName"`
		Plan                    string      `json:"plan"`
		PlanName                string      `json:"planName"`
		EnvironmentInstanceID   interface{} `json:"environmentInstanceId"`
		EnvironmentInstanceName interface{} `json:"environmentInstanceName"`
		SpaceID                 string      `json:"spaceId"`
		SpaceName               interface{} `json:"spaceName"`
		InstanceID              string      `json:"instanceId"`
		MeasureID               string      `json:"measureId"`
		MetricName              string      `json:"metricName"`
		UnitSingular            string      `json:"unitSingular"`
		UnitPlural              string      `json:"unitPlural"`
		IdentityZone            string      `json:"identityZone"`
		DataCenter              string      `json:"dataCenter"`
		DataCenterName          string      `json:"dataCenterName"`
		StartIsoDate            string      `json:"startIsoDate"`
		EndIsoDate              string      `json:"endIsoDate"`
		Usage                   float64     `json:"usage"`
	} `json:"content"`
}
