package client

import (
	"context"
	"gRPC/config"
	"gRPC/gRPC/paseto"
	auth "gRPC/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GRPCClient struct {
	prgcClient *grpc.ClientConn

	//proto 에서 생성된 authServiceClient
	authClient auth.AuthServiceClient

	// paseto 사용해야하니까 객체 심어두기
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

func (g *GRPCClient) CreateAuth(email string) (*auth.AuthData, error) {
	nowTime := time.Now()
	expiredTime := nowTime.Add(30 * time.Minute) // 만료시간 설정

	clientAuth := &auth.AuthData{
		Email:      email,
		CreateDate: nowTime.Unix(),
		ExpireDate: expiredTime.Unix(),
	}

	//paseto의 CreateNewToken 메소드 호출
	if token, err := g.pasetoMaker.CreateNewToken(clientAuth); err != nil {
		return nil, err
	} else {
		clientAuth.Token = token
		//서버에 있는 CreateAuth 를 타게됨 그래서 서버가 항상 켜져있어야해
		if response, err := g.authClient.CreateAuth(context.Background(), &auth.CreateTokenReq{Auth: clientAuth}); err != nil {
			return nil, err
		} else {
			return response.Auth, nil
		}
	}
}

func (g *GRPCClient) VerifyAuth(token string) (*auth.VerifyData, error) {
	if response, err := g.authClient.VerifyAuth(context.Background(), &auth.VerifyTokenReq{Token: token}); err != nil {
		return nil, err
	} else {
		return response.V, nil
	}
}
