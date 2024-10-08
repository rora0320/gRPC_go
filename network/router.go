// Package network 라우터 역할
// 라우터 -> 서비스를 통해 ->리포지토리에 요청-> 리포지토리 응답값을 서비스로 보내서 서비스에서 가공 -> 라우터로 가공 응답값 보냄
package network

import (
	"gRPC/config"
	"gRPC/gRPC/client"
	"gRPC/service"
	"github.com/gin-gonic/gin"
)

type Network struct {
	cfg     *config.Config
	service *service.Service

	// 프레임워크 같은것
	engin *gin.Engine

	//client 인증서버 -> 네트워크로 보내야해
	gRPCClient *client.GRPCClient
}

// NewNetwork 라우터->서비스
func NewNetwork(cfg *config.Config, service *service.Service, gRPCClient *client.GRPCClient) (*Network, error) {
	r := &Network{cfg: cfg, service: service, engin: gin.New(), gRPCClient: gRPCClient}

	//1. token 생성하는 api
	r.engin.POST("/login", r.login)
	//2. token 검증하는 api
	r.engin.GET("/verify", r.verifyLogin(), r.verify)
	return r, nil
}

// StartServer gin 이라는 http 웹 프레임워크 사용하기
func (s *Network) StartServer() {
	s.engin.Run(":8080")
}
