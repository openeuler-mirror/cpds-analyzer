package v1

import "github.com/emicklei/go-restful"

func AddToContainer(container *restful.Container) {
	webservice := new(restful.WebService)
	webservice.Path("/")

	handler := newRulesHandler()

	webservice.Route(webservice.GET("ping").
		To(handler.GetPing))

	container.Add(webservice)
}
