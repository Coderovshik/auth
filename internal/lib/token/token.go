package token

import (
	"time"

	"github.com/Coderovshik/auth/internal/domain"
	"github.com/golang-jwt/jwt/v5"
)

func New(user domain.User, app domain.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["exp"] = app.Info.ID

	return token.SignedString([]byte(app.Secret))
}
