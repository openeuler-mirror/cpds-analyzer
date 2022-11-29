package v1

import (
	"gitee.com/cpds/cpds-analyzer/pkgs/rules"
	"github.com/emicklei/go-restful"
)

type Handler struct {
	rules rules.Interface
}

func newRulesHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetRules(req *restful.Request, resp *restful.Response) {
	path := req.PathParameter("path")
	resp.WriteEntity(h.rules.LoadRules(path))
}
