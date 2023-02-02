package config

import (
	"fmt"

	"github.com/emicklei/go-restful"
	swagger "github.com/emicklei/go-restful-swagger12"
	"github.com/sirupsen/logrus"
)

const (
	apiPath = "/apidocs"
)

func (c *Config) RegisterSwagger(container *restful.Container) {
	logrus.Infof("registing swagger: http://%s:%s", c.BindAddress, c.Port)
	config := swagger.Config{
		WebServices:    container.RegisteredWebServices(),
		WebServicesUrl: fmt.Sprintf("http://%s:%s", c.BindAddress, c.Port),
		ApiPath:        apiPath,
	}
	swagger.RegisterSwaggerService(config, container)
}
