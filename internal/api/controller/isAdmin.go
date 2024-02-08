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

type IsAdmin struct {
	isAdminUsecase domain.IsAdimnUsecase
}

func NewIsAdmin(isAdminUsecase domain.IsAdimnUsecase) *IsAdmin {
	return &IsAdmin{
		isAdminUsecase: isAdminUsecase,
	}
}

func (i *IsAdmin) IsAdmin(ctx context.Context, req *desc.IsAdminRequest) (*desc.IsAdminResponse, error) {
	isAdmin, err := i.isAdminUsecase.IsAdmin(ctx, req.GetUserId())
	if err != nil {
		if errors.Is(err, usecase.ErrIdEmpty) {
			log.Printf("ERROR: failed to handle IsAdmin procedure call: %s\n", err.Error())

			return &desc.IsAdminResponse{
				IsAdmin: false,
			}, status.Error(codes.InvalidArgument, "user id empty")
		} else if errors.Is(err, usecase.ErrInvalidUserId) {
			log.Printf("ERROR: failed to handle IsAdmin procedure call: %s\n", err.Error())

			return &desc.IsAdminResponse{
				IsAdmin: false,
			}, status.Error(codes.InvalidArgument, "invalid user id")
		}

		log.Printf("ERROR: failed to handle IsAdmin procedure call: %s\n", err.Error())

		return &desc.IsAdminResponse{
			IsAdmin: false,
		}, status.Error(codes.Internal, "internal error")
	}

	return &desc.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
