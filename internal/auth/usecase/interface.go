package usecase

import "github.com/vera2005/lr11/internal/api"

type Provider interface {
	CheckUser(api.User) (api.User, error)
	CreateUser(api.User) error
	SelectUser(string) (api.User, error)
}
