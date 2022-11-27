package rules

import (
	"github.com/emicklei/go-restful"
)

type Rules struct {
	Cpu_max    int `json:"cpu_max"`
	Cpu_min    int `json:"cpu_min"`
	Disk_max   int `json:"disk_max"`
	Disk_min   int `json:"disk_min"`
	Memory_max int `json:"mem_max"`
	Memory_min int `json:"mem_min"`
}

func GetRules() *Rules {
	return &Rules{}
}

func (r *Rules) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/rules")

	container.Add(ws)
}
