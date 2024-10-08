// 리포지토리, 서비스, 네트워크에 대한 객체 값을 가지고 있음

package cmd

import (
	"gRPC/config"
	"gRPC/gRPC/client"
	"gRPC/network"
	"gRPC/repository"
	"gRPC/service"
)

type App struct {
	cfg *config.Config

	network    *network.Network
	repository *repository.Repository
	service    *service.Service

	gRPCClient *client.GRPCClient
}

// NewApp 리포지토리, 서비스, 네트워크에 대한 객체값이 필요해- 앱
func NewApp(cfg *config.Config) {

	var err error

	a := &App{cfg: cfg}

	client.NewGRPCClient(cfg)
	//3. http 라우터 client <->GRPC 서버 연결
	if a.gRPCClient, err = client.NewGRPCClient(cfg); err != nil {
		panic(err)

		//1. 리포지토리 - 디비연결 설정
	} else if a.repository, err = repository.NewRepository(cfg, a.gRPCClient); err != nil {
		panic(err)

		//2. 서비스 - 리포지토리로 연결
	} else if a.service, err = service.NewService(cfg, a.repository); err != nil {
		panic(err)
		//4. 네트워크(라우터) - 서비스로 연결
	} else if a.network, err = network.NewNetwork(cfg, a.service, a.gRPCClient); err != nil {
		panic(err)
	} else {
		//5. 서버 실행
		//TODO -> start server
		a.network.StartServer()
	}

}
