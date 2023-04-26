package errors

import "fmt"

const (
	SUCCESS        = 0
	DATABASE_ERROR = 101
	SOCKET_ERROR   = 102
	DETECTOR_ERROR = 103

	RULES_GET_ERROR    = 1001
	RULES_CREATE_ERROR = 1002
	RULES_UPDATE_ERROR = 1003
	RULES_DELETE_ERROR = 1004

	ANALYSIS_GET_RESULT_ERROR    = 2001
	ANALYSIS_DELETE_RESULT_ERROR = 2002
	ANALYSIS_GET_RAW_DATA_ERROR  = 2003

	MONITOR_GET_NODE_STATUS_ERROR              = 3001
	MONITOR_GET_NODE_INFO_ERROR                = 3002
	MONITOR_GET_NODE_RESOURCES_ERROR           = 3003
	MONITOR_GET_NODE_CONTAINER_STATUS_ERROR    = 3004
	MONITOR_GET_CLUSTER_RESOURCES_ERROR        = 3005
	MONITOR_GET_CLUSTER_CONTAINER_STATUS_ERROR = 3006
	MONITOR_GET_TARGET_ERROR                   = 3007
)

var AnalyzerResultCodeMap = map[uint16]string{
	SUCCESS:        "Success",
	DATABASE_ERROR: "Database Error",
	SOCKET_ERROR:   "Network Error",
	DETECTOR_ERROR: "Unable to connect to detector",

	RULES_GET_ERROR:    "Failed to get rule list",
	RULES_CREATE_ERROR: "Failed to create rule",
	RULES_UPDATE_ERROR: "Failed to update rule",
	RULES_DELETE_ERROR: "Failed to delete rule",

	ANALYSIS_GET_RESULT_ERROR:    "Failed to get analysis result",
	ANALYSIS_DELETE_RESULT_ERROR: "Failed to delete analysis result",
}

type Error struct {
	Err        error
	ResultCode uint16
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", AnalyzerResultCodeMap[e.ResultCode], e.Err.Error())
}

func NewError(resultCode uint16, err error) error {
	return &Error{
		ResultCode: resultCode,
		Err:        err,
	}
}

func IsErrorWithCode(err error, desiredResultCode uint16) bool {
	if err == nil {
		return false
	}

	serverError, ok := err.(*Error)
	if !ok {
		return false
	}

	return serverError.ResultCode == desiredResultCode
}
