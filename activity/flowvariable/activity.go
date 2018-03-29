package flowvariable

import (
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

// define variables
// inputs : {name, type, value}
type FlowVariableActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &FlowVariableActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *FlowVariableActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *FlowVariableActivity) Eval(context activity.Context) (done bool, err error) {

	return true, nil
}
