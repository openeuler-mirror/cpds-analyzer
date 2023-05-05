package prometheus

type queryParams struct {
	Query string `json:"query"`
	Time  int64  `json:"time"`
}

type queryRangeParams struct {
	Query      string `json:"query"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	StepSecond int64  `json:"step"`
}
