package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Coderovshik/auth/internal/domain"
	"github.com/Coderovshik/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var _ domain.RegisterUsecase = (*Register)(nil)

type Register struct {
	userRepo domain.UserRepository
}

func NewRegister(userRepo domain.UserRepository) *Register {
	return &Register{
		userRepo: userRepo,
	}
}

func (r *Register) Register(ctx context.Context, info domain.UserInfo) (string, error) {
	const op = "usecase.Register"

	err := validateUserInfo(info)
	if err != nil {
		log.Printf("ERROR: failed to validate user info: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("ERROR: failed to generate pass hash: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := r.userRepo.Save(ctx, info.Email, passHash)
	if err != nil {
		if errors.Is(err, repository.ErrUserExists) {
			log.Printf("ERROR: user exists: %s\n", err.Error())

			return "", fmt.Errorf("%s: %w", op, err)
		}

		log.Printf("ERROR: failed to save user info: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
