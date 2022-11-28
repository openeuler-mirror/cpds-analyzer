package rules

type Rules struct {
	Cpu_max    int `json:"cpu_max"`
	Cpu_min    int `json:"cpu_min"`
	Disk_max   int `json:"disk_max"`
	Disk_min   int `json:"disk_min"`
	Memory_max int `json:"mem_max"`
	Memory_min int `json:"mem_min"`
}

type Interface interface {
	LoadRules(path string) error
	SaveRules(path string) error
}

func New() *Rules {
	return &Rules{}
}
