package main

import (
	"fmt"

	"github.com/micro/go-plugins/config/source/grpc/v2"
	"github.com/zbrechave/tsquare/basic/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/cli/v2"
	"github.com/zbrechave/tsquare/srv/user/model"

	"github.com/micro/go-micro/v2/registry"
	"github.com/zbrechave/tsquare/basic/config"
	"github.com/zbrechave/tsquare/srv/user/handler"

	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/zbrechave/tsquare/basic"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	user "github.com/zbrechave/tsquare/srv/user/proto/user"
)

var (
	appName = "user_srv"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()
	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) error {
		// 初始化模型层
		model.Init()
		// 初始化handler
		handler.Init()

		return nil
	}))

	// Register Handler
	_ = user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)
	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Infof("[initCfg] 配置，cfg：%v", cfg)

	return
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}