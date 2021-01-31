package micro

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

// New 创建服务
func New(req *CreateMicroRequest) (*Micro, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	ins := &Micro{
		Id:              xid.New().String(),
		CreateAt:        ftime.Now().Timestamp(),
		UpdateAt:        ftime.Now().Timestamp(),
		Enabled:         true,
		Type:            req.Type,
		Name:            req.Name,
		Label:           req.Label,
		Description:     req.Description,
		TokenExpireTime: req.TokenExpireTime,
		RoleId:          req.RoleId,
	}

	return ins, nil
}

// NewCreateMicroRequest todo
func NewCreateMicroRequest() *CreateMicroRequest {
	return &CreateMicroRequest{
		Label: map[string]string{},
		Type:  Type_BUILD_IN,
	}
}

// Validate 校验请求是否合法
func (req *CreateMicroRequest) Validate() error {
	return validate.Struct(req)
}

// NewMicroSet 实例化
func NewMicroSet() *Set {
	return &Set{
		Items: []*Micro{},
	}
}

// Add 添加
func (s *Set) Add(e *Micro) {
	s.Items = append(s.Items, e)
}