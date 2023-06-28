package env

import (
	"cpea_monthly_usage/model"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

const ()

// get threshold values from env for services
func GetService() (map[string][]model.ServiceMetrics, error) {
	services := os.Getenv(SERVICES_ENV_NAME)

	if services != "" {

		var servicesThreshold map[string][]model.ServiceMetrics
		err := json.Unmarshal([]byte(services), &servicesThreshold)

		if err == nil {
			return servicesThreshold, nil
		} else {
			return nil, err
		}

	}

	return Services_Value_Map_Constants, nil
}

// get UAS config
func GetUASConfigValue() (model.UASConfig, error) {

	uasConfig := os.Getenv(UAS_ENV_NAME)
	var config model.UASConfig
	if uasConfig != "" {

		err := json.Unmarshal([]byte(uasConfig), &config)

		if err == nil {
			return config, nil
		}

	}

	return config, errors.New("error getting uas values")

}

// get Dest config
func GetDestConfigValue() (model.DestConfig, error) {

	destConfig := os.Getenv(Dest_ENV_NAME)
	var config model.DestConfig
	if destConfig != "" {

		err := json.Unmarshal([]byte(destConfig), &config)

		if err == nil {
			return config, nil
		}

	}

	return config, errors.New("error getting destination env values")

}

// get time threshold in minutes
func GetTimeInMinutes() time.Duration {

	t := os.Getenv(TIME_ENV_NAME)

	if t != "" {
		t, err := strconv.Atoi(t)
		if err == nil {
			return time.Duration(t) * time.Minute
		} else {
			return Time_Constant * time.Minute
		}
	}
	return Time_Constant * time.Minute

}

// get Env for Event name

func GetEventName() string {

	event := os.Getenv(Event_TYPE_ENV_NAME)

	if event != "" {
		return event
	}
	return EventName
}

// get Env For Sub account mode

func GetSubaccountMode() bool {
	subaccountMode := os.Getenv(MODE_SUBACCOUNT_ENV_NAME)

	if subaccountMode != "" {
		subaccountBool, err := strconv.ParseBool(subaccountMode)
		if err != nil {
			return false
		}
		return subaccountBool
	}
	return SubaccountMode
}
