package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/tag"
	"github.com/infraboard/keyauth/pkg/token"
)

func newQueryTagKeyRequest(tk *token.Token, req *tag.QueryTagKeyRequest) (*queryTagKeyRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTagKeyRequest{
		tk:                 tk,
		QueryTagKeyRequest: req}, nil
}

type queryTagKeyRequest struct {
	tk *token.Token
	*tag.QueryTagKeyRequest
}

func (r *queryTagKeyRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTagKeyRequest) FindFilter() bson.M {
	filter := bson.M{}

	// if r.Type != role.RoleType_NULL {
	// 	filter["type"] = r.Type
	// } else {
	// 	// 获取内建和全局的角色以及域内自己创建的角色
	// 	filter["$or"] = bson.A{
	// 		bson.M{"type": role.RoleType_BUILDIN},
	// 		bson.M{"type": role.RoleType_GLOBAL},
	// 		bson.M{"type": role.RoleType_CUSTOM, "domain": r.tk.Domain},
	// 	}
	// }

	return filter
}

func newQueryTagValueRequest(tk *token.Token, req *tag.QueryTagValueRequest) (*queryTagValueRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return &queryTagValueRequest{
		tk:                   tk,
		QueryTagValueRequest: req}, nil
}

type queryTagValueRequest struct {
	tk *token.Token
	*tag.QueryTagValueRequest
}

func (r *queryTagValueRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryTagValueRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.TagId != "" {
		filter["key_id"] = r.TagId
	}

	return filter
}