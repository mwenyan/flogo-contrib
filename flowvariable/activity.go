package flowvariable

import (
	"fmt"
	"strconv"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-flowvariable")

const (
	ivName   = "name"
	ivValue  = "value"
	ivType = "type"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// create or update state
// inputs : {assetName, assetValue}
// outputs: asset
type SmartContractActivity struct {
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

	name, _ := context.GetInput(ivName).(string)
	_type, _ := context.GetInput(ivType).(string)
	value, _ := context.GetInput(ivValue).(string)

	return true, nil
}