package processing

import "fmt"

func process() {

	// get threshold Values
	thresholds := getThresholdResponse()
	fmt.Println("Threshold Values After Calcukation", thresholds)

	// check if usage has exceeded and trigger notification
	checkThresholdAndtriggerAlert(thresholds)

}

func processWithSubaccount() {

	// get threshold Values
	thresholds := getThresholdResponseWithSubaccount()
	fmt.Println("Threshold Values After Calcukation", thresholds)

	// check if usage has exceeded and trigger notification
	checkThresholdAndtriggerAlertWithSubaccount(thresholds)

}
