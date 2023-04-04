package processing

import (
	"cpea_monthly_usage/env"
	"os"
	"time"
)

var processFunc = process
var processFuncSubaccount = processWithSubaccount
var timeVal = env.GetTimeInMinutes()
var endChannel = env.EndChan
var osexit = os.Exit

func CheckMonthlyData() {

	// timeVal := time.Duration(env.GetTimeInMinutes()) * time.Second
	ticker := time.NewTicker(timeVal)

	go func() {
		for {
			select {
			case <-ticker.C:
				{

					// Calling the Process
					if !env.GetSubaccountMode() {
						processFunc()
					} else {
						processFuncSubaccount()
					}

				}

			}
		}
	}()
}

func CheckErrorAndEndApplication() {
	for {
		select {
		case <-endChannel:
			{
				osexit(1)
			}
		}
	}

}
