go test -coverprofile cover.out ./api/... ./processing/... ./env/...
go tool cover -html=cover.out 

`Variables required in Manifest`

* Time in Minute
* UAS Service Credentials
* Values To Be Put In Threshold

# Format of Variables
`Time For App to Scrape Metrics`
```
5
```

`Event type For Alert Notification`
```
cpeacreditsover
```


`Destination Config`

```
{
    "Client_Id":"<client_id>",
    "Client_Secret":"<client_secret>",
    "Token_Url":"<token_url>",
    "Url":"<url>",
    "DestName":"<destination_name"
}
```

`UAS Credential Service`
```
{
    "Url":"<url_of_service>",
    "Client_Id":"<client_id>",
    "Client_Secret":"<client_secret>",
    "Token_Url":"<Token_url>"
    
}
```
`Threshold Values - Mode 1`
`SubaccountMode - false `
```
{
    "<Service_ID>":[{
        "Metric":"<measureId>",
        "PlanId":"<plan>",
        "Threshold":"<threshold>"
    }]
}
```
# e.g.
```
{
  "CloudIntegration":[
    {"Metric":"connections","ThreholdValue":100,"PlanId":"standard"},
    {"Metric":"tenants","ThreholdValue":5,"PlanId":"standard"}
    ],
    "sap-workzone":[
    {"Metric":"swz_users","ThreholdValue":10,"PlanId":"standard"},
    {"Metric":"swz_connections","ThreholdValue":10,"PlanId":"standard"}
    ]
}

```
`Threshold Values - Mode 2`
`SubaccountMode - true `
```
{
    "<Service_ID>":[{
        "Metric":"<measureId>",
        "PlanId":"<plan>",
        "Threshold":"<threshold>",
        "SubaccountId":"<subaccountId>"
    }]
}
```
# e.g.
```
{
  "CloudIntegration":[
    {"Metric":"connections","ThreholdValue":100,"PlanId":"standard", "SubaccountId": "87e47fc4-b65c-4891"},
    {"Metric":"tenants","ThreholdValue":5,"PlanId":"standard","SubaccountId": "87e47fc4-b65c-4891"}
    ],
    "sap-workzone":[
    {"Metric":"swz_users","ThreholdValue":10,"PlanId":"standard","SubaccountId": "87e47fc4-b65c-4891"},
    {"Metric":"swz_connections","ThreholdValue":10,"PlanId":"standard","SubaccountId": "87e47fc4-b65c-4891"}
    ]
}
```