package memory

import (
	"go-sandbox/hexagonal/shortener"
	"time"
)

type Repository struct{}

var redirects = []shortener.Redirect{
	{Code: "abcd", Url: "http://www.google.com", CreatedAt: time.Now().Unix()},
	{Code: "efgh", Url: "http://www.stackoverflow.com", CreatedAt: time.Now().Unix()},
}

func (r Repository) Find(code string) (shortener.Redirect, error) {
	for _, redirect := range redirects {
		if redirect.Code == code {
			return redirect, nil
		}
	}
	return shortener.Redirect{}, shortener.ErrRedirectNotFound
}

func (r Repository) Store(redirect shortener.Redirect) error {
	redirects = append(redirects, redirect)
	return nil
}

func (r Repository) FindAll() []shortener.Redirect {
	return redirects
}
