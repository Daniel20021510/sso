package jwt

import (
	"github.com/Daniel20021510/sso/internal/domain/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user *model.User, app *model.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
