package middleware

import (
	"os"
	"shopBackend/app/repository"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserMiddleware struct {
	repo repository.UserRepoInterface
}

func NewUserMiddleware(repo repository.UserRepoInterface) *UserMiddleware {
	return &UserMiddleware{repo}
}

func CreateToken(id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
