package rest

import (
	"github.com/James-Milligan/FutureNetworksBU/internal/common"
	"github.com/goioc/di"
	"reflect"
)

func BuildDependencyContainer() {

	_, _ = di.RegisterBeanInstance("config", common.GetConfig())
	_, _ = di.RegisterBean("grpcClient", reflect.TypeOf((*GrpcClient)(nil)))

	_, _ = di.RegisterBean("app", reflect.TypeOf((*App)(nil)))

	_ = di.InitializeContainer()
}
