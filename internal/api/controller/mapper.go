package controller

import (
	"github.com/Coderovshik/auth/internal/domain"
	desc "github.com/Coderovshik/auth/pkg/auth"
)

func toDomainUserInfo(info *desc.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Email:    info.GetEmail(),
		Password: info.GetPassword(),
	}
}

func toDomainAppInfo(info *desc.AppInfo) domain.AppInfo {
	return domain.AppInfo{
		ID: info.GetId(),
	}
}
