package controller

import (
	"context"
	"errors"
	"log"

	"github.com/Coderovshik/auth/internal/domain"
	"github.com/Coderovshik/auth/internal/usecase"
	desc "github.com/Coderovshik/auth/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Login struct {
	loginUsecase domain.LoginUsecase
}

func NewLogin(loginUsecase domain.LoginUsecase) *Login {
	return &Login{
		loginUsecase: loginUsecase,
	}
}

func (lc *Login) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	token, err := lc.loginUsecase.Login(
		ctx,
		toDomainUserInfo(req.GetUserInfo()),
		toDomainAppInfo(req.GetAppInfo()),
	)
	if err != nil {
		if errors.Is(err, usecase.ErrEmailEmpty) {
			log.Printf("ERROR: failed to handle Login procedure call: %s\n", err.Error())

			return &desc.LoginResponse{
				Token: "",
			}, status.Error(codes.InvalidArgument, "empty email")
		} else if errors.Is(err, usecase.ErrPasswordEmpty) {
			log.Printf("ERROR: failed to handle Login procedure call: %s\n", err.Error())

			return &desc.LoginResponse{
				Token: "",
			}, status.Error(codes.InvalidArgument, "empty password")
		} else if errors.Is(err, usecase.ErrInvalidCredentials) {
			log.Printf("ERROR: failed to handle Login procedure call: %s\n", err.Error())

			return &desc.LoginResponse{
				Token: "",
			}, status.Error(codes.InvalidArgument, "invalid credentials")
		} else if errors.Is(err, usecase.ErrInvalidAppInfo) {
			log.Printf("ERROR: failed to handle Login procedure call: %s\n", err.Error())

			return &desc.LoginResponse{
				Token: "",
			}, status.Error(codes.InvalidArgument, "invalid app info")
		}

		log.Printf("ERROR: failed to handle Login procedure call: %s\n", err.Error())

		return &desc.LoginResponse{
			Token: "",
		}, status.Error(codes.Internal, "internal error")
	}

	return &desc.LoginResponse{
		Token: token,
	}, nil
}
