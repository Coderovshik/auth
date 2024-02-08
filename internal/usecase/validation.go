package usecase

import "github.com/Coderovshik/auth/internal/domain"

func validateUserInfo(info domain.UserInfo) error {
	if len(info.Email) == 0 {
		return ErrEmailEmpty
	}

	if len(info.Password) == 0 {
		return ErrPasswordEmpty
	}

	return nil
}

func validateId(id string) error {
	if len(id) == 0 {
		return ErrIdEmpty
	}

	return nil
}
