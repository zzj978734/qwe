package http

import (
	"errors"

	"github.com/infraboard/mcube/http/router"

	"github.com/infraboard/keyauth/client"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/session"
)

var (
	api = &handler{}
)

type handler struct {
	service session.UserServiceClient
}

// Registry 注册HTTP服务路由
func (h *handler) Registry(router router.SubRouter) {
	r := router.ResourceRouter("sessions")
	r.BasePath("sessions")
	r.Permission(true)
	r.Handle("GET", "/", h.QueryLoginLog)
}

func (h *handler) Config() error {
	client := client.C()
	if client == nil {
		return errors.New("grpc client not initial")
	}

	h.service = client.SessionUser()
	return nil
}

func init() {
	pkg.RegistryHTTPV1("session", api)
}
