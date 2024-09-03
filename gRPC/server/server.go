package server

import (
	"gRPC/config"
	"gRPC/gRPC/paseto"
	auth "gRPC/gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPCServer struct {
	pasetoMaker    paseto.PasetoMaker
	tokenVerifyMap map[string]auth.AuthData
}

func NewGRPCServer(cfg *config.Config) error {
	if lis, err := net.Listen("tcp", cfg.GRPC.URL); err != nil {
		return err
	} else {
		server := grpc.NewServer([]grpc.ServerOption{}...)
		auth.RegisterAuthServiceServer(server, &GRPCServer{
			pasetoMaker: paseto.NewPasetoMaker(cfg),
			tokenVerifyMap: make(map[string]*auth.AuthData)
		})

		//server에 반영서비스를 등록한다
		reflection.Register(server)

		//서버가 실행이 되면 그이후의 코드가 실행되지 않아서 스레드를 따로 생성해 실행하도록 함.
		//스레드 생성 -> 백그라운드에서 돌린다고 생각하면 됨.
		go func(){
			//server.Serve() 서버가 실행 된 이후의 코드는 실행이 되지않음!
			if err := server.Serve(lis); err != nil {
				panic(err)
			}
		}()

		return nil
	}
}
