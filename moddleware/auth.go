package middleware

import "shopBackend/app/repository"

type UserMiddleware struct {
	repo repository.UserRepoInterface
}

func NewUserMiddleware(repo repository.UserRepoInterface) *UserMiddleware {
	return &UserMiddleware{repo}
}
