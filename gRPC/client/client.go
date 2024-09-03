package client

import (
	"gRPC/config"
	"gRPC/gRPC/paseto"
	auth "gRPC/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	prgcClient *grpc.ClientConn
	authClient auth.AuthServiceClient

	// paseto 사용해야하니까
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(cfg *config.Config) (*GRPCClient, error) {
	c := new(GRPCClient)

	// prgc 패키지 내 WithTransportCredentials
	if prgcClient, err := grpc.NewClient(cfg.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	} else {
		c.prgcClient = prgcClient
		c.authClient = auth.NewAuthServiceClient(c.prgcClient)
		c.pasetoMaker = paseto.NewPasetoMaker(cfg)
	}
	return c, nil
}

//rpc CreateAuth(CreateTokenReq) returns (CreateTokenRes);
//rpc VerifyAuth(CreateTokenReq) returns (CreateTokenRes);

func (g *GRPCClient) CreateAuth(address string) (*auth.AuthData, error) {
	return nil, nil
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.VerifyTokenRes, error) {
	return nil, nil
}
