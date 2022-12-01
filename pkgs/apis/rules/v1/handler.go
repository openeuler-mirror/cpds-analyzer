package v1

import (
	"encoding/json"
	"io"

	"gitee.com/cpds/cpds-analyzer/pkgs/rules"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	rules rules.Rules
}

func newRulesHandler(r *rules.Rules) *Handler {
	return &Handler{
		rules: *r,
	}
}

func (h *Handler) GetRules(req *restful.Request, resp *restful.Response) {
	path := req.QueryParameter("path")

	if path == "" {
		path = h.rules.GetDefaultRulesPath()
	}

	if err := h.rules.LoadRules(path); err != nil {
		logrus.Error(err)
	}

	b, _ := json.Marshal(h.rules)
	io.WriteString(resp.ResponseWriter, string(b))
}
