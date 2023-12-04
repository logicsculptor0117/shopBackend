package middleware

import (
	"fmt"
	"os"
	"shopBackend/app/repository"
	"shopBackend/helpers"
	"time"

	"github.com/gin-gonic/gin"
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

func (um UserMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get token
		token, err := ValidateToken(ctx.Request.Header.Get("Authorization"))
		if err != nil {
			fError := helpers.FieldError{Field: "token", Message: err.Error()}
			helpers.RespondJSON(ctx, 401, helpers.StatusCodeFromInt(401), fError, nil)
			ctx.Abort()
			return
		}
		user, errDB := um.repo.GetUserFromToken(token)
		if errDB != nil {
			status, errorDB := helpers.DBError(errDB)
			helpers.RespondJSON(ctx, status, helpers.StatusCodeFromInt(status), errorDB, nil)
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	//Parse token string into a token object
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method and return the secret key
		if token.Method != token.Method.(*jwt.SigningMethodHMAC) {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_JWT")), nil
	})
}
