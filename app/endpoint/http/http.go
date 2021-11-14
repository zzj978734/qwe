package http

import (
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/app/endpoint"
	"github.com/infraboard/keyauth/app/user/types"
)

var (
	api = &handler{}
)

type handler struct {
	endpoint endpoint.ServiceServer
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("endpoint")

	r.BasePath("endpoints")
	r.Handle("POST", "/", h.Create).SetAllow(types.UserType_INTERNAL)
	r.Handle("GET", "/", h.List)
	r.Handle("GET", "/:id", h.Get)

	rr := router.ResourceRouter("resource")
	rr.BasePath("resources")
	rr.Handle("GET", "/", h.ListResource)
}

func (h *handler) Config() error {
	h.endpoint = app.GetGrpcApp(endpoint.AppName).(endpoint.ServiceServer)
	return nil
}
