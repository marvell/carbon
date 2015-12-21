package carbon

import (
	"fmt"

	"github.com/marvell/carbon/loggers"
	"github.com/marvell/carbon/servers"
)

func NewApi(name, version string) *Api {
	api := &Api{
		Name:    name,
		Version: version,
	}

	// set defaults
	api.SetLogger(loggers.NewConsoleLogger())
	api.SetHttpServer(servers.NewHttp1Server())

	return api
}

type Api struct {
	Name      string
	Version   string
	UrlPrefix string

	Logger     Logger
	HttpServer HttpServer
}

func (api *Api) SetLogger(l Logger) {
	api.Logger = l
	api.Logger.SetPrefix(fmt.Sprintf("[%s] ", api.Name))
}

func (api *Api) SetHttpServer(s HttpServer) {
	api.HttpServer = s
}

func (api *Api) RunHttpServer() {
	api.Logger.Infof("Starting http-server on %s", api.HttpServer.BindAddress())

	err := api.HttpServer.Run()
	if err != nil {
		api.Logger.Error(err)
	}
}
