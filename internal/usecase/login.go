package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Coderovshik/auth/internal/domain"
	"github.com/Coderovshik/auth/internal/lib/token"
	"github.com/Coderovshik/auth/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var _ domain.LoginUsecase = (*Login)(nil)

type Login struct {
	userRepo domain.UserRepository
	appRepo  domain.AppRepository
	tokenTTL time.Duration
}

func NewLogin(userRepo domain.UserRepository, appRepo domain.AppRepository, tokenTTL time.Duration) *Login {
	return &Login{
		userRepo: userRepo,
		appRepo:  appRepo,
		tokenTTL: tokenTTL,
	}
}

func (l *Login) Login(ctx context.Context, userInfo domain.UserInfo, appInfo domain.AppInfo) (string, error) {
	const op = "usecase.Login"

	err := validateUserInfo(userInfo)
	if err != nil {
		log.Printf("ERROR: failed to validate user info: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	user, err := l.userRepo.Get(ctx, userInfo.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			log.Printf("ERROR: user not found: %s\n", err.Error())

			return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		}

		log.Printf("ERROR: failed to get user from repository: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(userInfo.Password)); err != nil {
		log.Printf("ERROR: hash and password comparison failed: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	app, err := l.appRepo.Get(ctx, appInfo.ID)
	if err != nil {
		if errors.Is(err, repository.ErrAppNotFound) {
			log.Printf("ERROR: app not found: %s\n", err.Error())

			return "", fmt.Errorf("%s: %w", op, ErrInvalidAppInfo)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	log.Print("user logged in successfully")

	tkn, err := token.New(user, app, l.tokenTTL)
	if err != nil {
		log.Printf("ERROR: failed to generate token: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return tkn, nil
}
