package start

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-start")

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}


type StartActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SmartContractActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *SmartContractActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *SmartContractActivity) Eval(context activity.Context) (done bool, err error) {

	return true, nil
}