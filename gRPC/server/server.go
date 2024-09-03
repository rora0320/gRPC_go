package server

import (
	"context"
	"gRPC/config"
	"gRPC/gRPC/paseto"
	auth "gRPC/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type GRPCServer struct {
	auth.AuthServiceServer
	pasetoMaker    *paseto.PasetoMaker
	tokenVerifyMap map[string]*auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)
		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker:    paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData),
		})

		//server에 반영서비스를 등록한다
		reflection.Register(server)

		//서버가 실행이 되면 그이후의 코드가 실행되지 않아서 스레드를 따로 생성해 실행하도록 함.
		//스레드 생성 -> 백그라운드에서 돌린다고 생각하면 됨.
		go func() {
			//server.Serve() 서버가 실행 된 이후의 코드는 실행이 되지않음!
			if err := server.Serve(lis); err != nil {
				panic(err)
			}
		}()

		return nil
	}
}

func (s *GRPCServer) CreateAuth(_ context.Context, req *auth.CreateTokenReq) (*auth.CreateTokenRes, error) {
	data := req.Auth
	token := data.Token
	s.tokenVerifyMap[token] = data

	return &auth.CreateTokenRes{Auth: data}, nil
}

func (s *GRPCServer) VerifyAuth(_ context.Context, req *auth.VerifyTokenReq) (*auth.VerifyTokenRes, error) {
	token := req.Token

	//토큰 검증 로직 호출 res 받아서
	res := &auth.VerifyTokenRes{V: &auth.VerifyData{
		Auth: nil,
	}}

	//응답의 status 업데이트
	if authData, ok := s.tokenVerifyMap[token]; !ok {
		res.V.Status = auth.ResponseType_FAILED
	} else if authData.ExpireDate < time.Now().Unix() {
		delete(s.tokenVerifyMap, token)
		res.V.Status = auth.ResponseType_EXPIRED_DATE
	} else {
		res.V.Status = auth.ResponseType_SUCCESS
	}
	return res, nil
}
