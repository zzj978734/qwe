package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"github.com/infraboard/keyauth/conf"
	"github.com/infraboard/keyauth/pkg"
	"github.com/infraboard/keyauth/pkg/namespace"
	"github.com/infraboard/keyauth/pkg/permission"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/role"
	"github.com/infraboard/keyauth/pkg/user"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	log           logger.Logger
	col           *mongo.Collection
	enableCache   bool
	notifyCachPre string

	namespace namespace.NamespaceServiceServer
	user      user.UserServiceServer
	role      role.RoleServiceServer

	policy.UnimplementedPolicyServiceServer
}

func (s *service) Config() error {
	if pkg.Namespace == nil {
		return fmt.Errorf("dependence namespace service is nil, please load first")
	}
	s.namespace = pkg.Namespace

	if pkg.User == nil {
		return fmt.Errorf("dependence user service is nil, please load first")
	}
	s.user = pkg.User

	if pkg.Role == nil {
		return fmt.Errorf("dependence role service is nil, please load first")
	}
	s.role = pkg.Role

	db := conf.C().Mongo.GetDB()
	col := db.Collection("policy")

	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{{Key: "create_at", Value: bsonx.Int32(-1)}},
		},
	}

	_, err := col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}

	s.col = col
	s.log = zap.L().Named("Policy")
	return nil
}

func (s *service) Name() string {
	return permission.AppName
}

func (s *service) Registry(server *grpc.Server) {
	policy.RegisterPolicyServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}