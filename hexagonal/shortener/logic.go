package shortener

import (
	"errors"
	"go-sandbox/hexagonal/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate = validator.New()

var (
	ErrRedirectNotFound = errors.New("redirect not found")
	ErrRedirectInvalid  = errors.New("invalid url format")
)

type service struct {
	repo Repository
}

func NewService(repo Repository) service {
	return service{repo}
}

func (s service) Find(code string) (Redirect, error) {
	return s.repo.Find(code)
}

func (s service) Create(url string) (Redirect, error) {
	redirect := Redirect{
		Url:       url,
		Code:      utils.RandString(4),
		CreatedAt: time.Now().Unix(),
	}

	if err := validate.Struct(redirect); err != nil {
		return Redirect{}, ErrRedirectInvalid
	}

	if err := s.repo.Store(redirect); err != nil {
		return Redirect{}, err
	}

	return redirect, nil
}

func (s service) FindAll() []Redirect {
	return s.repo.FindAll()
}
