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

type Register struct {
	registerUsecase domain.RegisterUsecase
}

func NewRegister(registerUsecase domain.RegisterUsecase) *Register {
	return &Register{
		registerUsecase: registerUsecase,
	}
}

func (rc *Register) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {
	id, err := rc.registerUsecase.Register(ctx, toDomainUserInfo(req.GetInfo()))
	if err != nil {
		if errors.Is(err, usecase.ErrEmailEmpty) {
			log.Printf("ERROR: failed to handle Register procedure call: %s\n", err.Error())

			return &desc.RegisterResponse{
				UserId: "",
			}, status.Error(codes.InvalidArgument, "email empty")
		} else if errors.Is(err, usecase.ErrPasswordEmpty) {
			log.Printf("ERROR: failed to handle Register procedure call: %s\n", err.Error())

			return &desc.RegisterResponse{
				UserId: "",
			}, status.Error(codes.InvalidArgument, "password empty")
		} else if errors.Is(err, usecase.ErrInvalidCredentials) {
			log.Printf("ERROR: failed to handle Register procedure call: %s\n", err.Error())

			return &desc.RegisterResponse{
				UserId: "",
			}, status.Error(codes.InvalidArgument, "invalid credentials")
		}

		log.Printf("ERROR: failed to handle Register procedure call: %s\n", err.Error())

		return &desc.RegisterResponse{
			UserId: "",
		}, status.Error(codes.Internal, "internal error")
	}

	// TODO: make proper error processing

	return &desc.RegisterResponse{
		UserId: id,
	}, nil
}
