package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Coderovshik/auth/internal/domain"
	"github.com/Coderovshik/auth/internal/repository"
)

var _ domain.IsAdimnUsecase = (*IsAdmin)(nil)

type IsAdmin struct {
	userRepo domain.UserRepository
}

func NewIsAdmin(userRepo domain.UserRepository) *IsAdmin {
	return &IsAdmin{
		userRepo: userRepo,
	}
}

func (i *IsAdmin) IsAdmin(ctx context.Context, id string) (bool, error) {
	const op = "usecase.IsAdmin"

	err := validateId(id)
	if err != nil {
		log.Printf("ERROR: failed to validate user id: %s\n", err.Error())

		return false, fmt.Errorf("%s: %w", op, err)
	}

	isAdmin, err := i.userRepo.IsAdmin(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Printf("ERROR: user not found: %s\n", err.Error())

			return false, fmt.Errorf("%s: %w", op, ErrInvalidUserId)
		}

		log.Printf("ERROR: failed to get user admin status: %s\n", err.Error())

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
