package v1

import (
	"io"

	"github.com/emicklei/go-restful"
)

type Handler struct {
}

func newRulesHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetPing(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "success")
}
