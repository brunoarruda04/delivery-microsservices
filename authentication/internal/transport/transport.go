package transport

import (
	"authentication/internal/domain"
	"authentication/internal/service"
	"authentication/proto"
	"context"
)

type grpcServer struct {
	svc service.AuthService
	proto.UnimplementedAuthServiceServer
}

func NewGRPCServer(svc service.AuthService) proto.AuthServiceServer {
	return &grpcServer{svc: svc}
}

func (s *grpcServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	id, err := s.svc.Register(req.Username, req.Password, domain.Role(req.Role))
	if err != nil {
		return nil, err
	}
	return &proto.RegisterResponse{Id: id}, nil
}

func (s *grpcServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	id, err := s.svc.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{Id: id}, nil
}
