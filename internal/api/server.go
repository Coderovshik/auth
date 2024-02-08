package api

import (
	"context"

	"github.com/Coderovshik/auth/internal/api/controller"
	desc "github.com/Coderovshik/auth/pkg/auth"
)

var _ desc.AuthServer = (*Server)(nil)

type Server struct {
	desc.UnimplementedAuthServer
	loginController    *controller.Login
	isAdminController  *controller.IsAdmin
	registerController *controller.Register
}

func NewServer(lc *controller.Login, ic *controller.IsAdmin, rc *controller.Register) *Server {
	return &Server{
		UnimplementedAuthServer: desc.UnimplementedAuthServer{},
		loginController:         lc,
		isAdminController:       ic,
		registerController:      rc,
	}
}

func (s *Server) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	return s.registerController.Register(ctx, req)
}

func (s *Server) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	return s.loginController.Login(ctx, req)
}

func (s *Server) IsAdmin(ctx context.Context, req *desc.IsAdminRequest) (*desc.IsAdminResponse, error) {
	return s.isAdminController.IsAdmin(ctx, req)
}
