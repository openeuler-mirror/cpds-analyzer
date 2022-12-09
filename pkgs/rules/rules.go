package rules

import "gitee.com/cpds/cpds-analyzer/utils"

type Rules struct {
	Cpu_max    int `json:"cpu_max"`
	Cpu_min    int `json:"cpu_min"`
	Disk_max   int `json:"disk_max"`
	Disk_min   int `json:"disk_min"`
	Memory_max int `json:"mem_max"`
	Memory_min int `json:"mem_min"`
}

const (
	defaultRulesPath = "/etc/cpds/cpds-analyzer/rules.json"
)

func New() *Rules {
	return &Rules{}
}

func (r *Rules) LoadRules(path string) error {
	if err := utils.LoadJsonFromFile(path, r); err != nil {
		return err
	}
	return nil
}

func (r *Rules) SetRules(path string) error {
	if err := utils.SaveAsJsonFile(path, r); err != nil {
		return err
	}
	return nil
}

func (r *Rules) GetDefaultRulesPath() string {
	return defaultRulesPath
}
