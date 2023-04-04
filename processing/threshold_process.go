package processing

import (
	"cpea_monthly_usage/env"
	"cpea_monthly_usage/model"
	"fmt"
)

var EndChan = env.EndChan

func getThresholdResponse() map[[3]string]model.MetricsThrehsoldValue {

	uasResponse := getUASResponse()
	thresholds := checkThreshold(uasResponse)

	return thresholds

}

func getThresholdResponseWithSubaccount() map[[4]string]model.MetricsThrehsoldValue {

	uasResponse := getUASResponse()
	thresholds := checkThresholdWithSubaccount(uasResponse)

	return thresholds

}

func checkThreshold(uasResponse model.UASMonthlyUsageResponse) map[[3]string]model.MetricsThrehsoldValue {

	metricsThresholdValue := map[[3]string]model.MetricsThrehsoldValue{}

	// Getting threshold values from ENV

	uasEnvValuesThreshold, err := env.GetService()

	// if Env Variables are not mapped properly
	if err != nil {
		fmt.Println("Error in getting env varible Service", err)
		EndChan <- 1
		return nil
	}

	for _, v := range uasResponse.Content {

		// check if Service is to be used for calculation
		if _, ok := uasEnvValuesThreshold[v.ServiceID]; ok {
			// check if metrics is there for threshold
			for _, val := range uasEnvValuesThreshold[v.ServiceID] {
				// if metrics && plan id match
				if val.Metric == v.MeasureID && val.PlanId == v.Plan {

					// check if metrics exist in the map else create one
					if _, ok := metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}]; ok {
						metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}] = model.MetricsThrehsoldValue{
							metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}].Value + int(v.Usage),
							metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}].PassedThrehold,
						}
					} else {
						metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}] = model.MetricsThrehsoldValue{
							int(v.Usage),
							false,
						}
					}

					// if measure if bigger then threshold
					if metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}].Value > val.ThreholdValue {
						metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}] = model.MetricsThrehsoldValue{
							metricsThresholdValue[[3]string{v.ServiceID, v.MeasureID, v.Plan}].Value,
							true,
						}
					}

					// once the macth happens break
					break
				}
			}
		}

	}

	return metricsThresholdValue

}

func checkThresholdWithSubaccount(uasResponse model.UASMonthlyUsageResponse) map[[4]string]model.MetricsThrehsoldValue {

	metricsThresholdValue := map[[4]string]model.MetricsThrehsoldValue{}

	// Getting threshold values from ENV

	uasEnvValuesThreshold, err := env.GetService()

	// if Env Variables are not mapped properly
	if err != nil {
		fmt.Println("Error in getting env varible Service", err)
		env.EndChan <- 1
	}

	for _, v := range uasResponse.Content {

		// check if Service is to be used for calculation
		if _, ok := uasEnvValuesThreshold[v.ServiceID]; ok {
			// check if metrics is there for threshold
			for _, val := range uasEnvValuesThreshold[v.ServiceID] {
				// if metrics, plan id and sub account id match
				if val.Metric == v.MeasureID && val.SubaccountId == v.SubaccountID && val.PlanId == v.Plan {

					// check if metrics exist in the map else create one
					if _, ok := metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}]; ok {
						metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}] = model.MetricsThrehsoldValue{
							metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}].Value + int(v.Usage),
							metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}].PassedThrehold,
						}
					} else {
						metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}] = model.MetricsThrehsoldValue{
							int(v.Usage),
							false,
						}
					}

					// if measure if bigger then threshold
					if metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}].Value > val.ThreholdValue {
						metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}] = model.MetricsThrehsoldValue{
							metricsThresholdValue[[4]string{v.ServiceID, v.MeasureID, v.Plan, v.SubaccountID}].Value,
							true,
						}
					}

					// once the macth happens break
					break
				}

			}
		}

	}

	return metricsThresholdValue

}
