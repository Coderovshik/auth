package domain

import "context"

type LoginUsecase interface {
	Login(ctx context.Context, userInfo UserInfo, appInfo AppInfo) (string, error)
}
