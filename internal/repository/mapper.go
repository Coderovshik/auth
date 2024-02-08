package repository

import "github.com/Coderovshik/auth/internal/domain"

func toDomainUser(doc userDocument) domain.User {
	return domain.User{
		ID:       doc.ID.Hex(),
		Email:    doc.Email,
		PassHash: doc.PassHash,
	}
}

func toDomainApp(doc appDocument) domain.App {
	return domain.App{
		Info: domain.AppInfo{
			ID: doc.ID.Hex(),
		},
		Name:   doc.Name,
		Secret: doc.Secret,
	}
}
