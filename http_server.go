package carbon

import (
	"net/http"
	"time"
)

type HttpServer interface {
	BindAddress() string
	SetBindAddress(string)

	SetTimeouts(time.Duration, time.Duration)
	SetDispatcher(http.Handler)

	Run() error
}
