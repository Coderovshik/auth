package domain

import "context"

type IsAdimnUsecase interface {
	IsAdmin(ctx context.Context, id string) (bool, error)
}
