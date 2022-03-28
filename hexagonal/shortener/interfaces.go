package shortener

type Service interface {
	Find(code string) (Redirect, error)
	Create(url string) (Redirect, error)
	FindAll() []Redirect
}

type Repository interface {
	Find(code string) (Redirect, error)
	Store(redirect Redirect) error
	FindAll() []Redirect
}

type Serializer interface {
	Encode(input Redirect) ([]byte, error)
	Decode(input []byte) (Redirect, error)
}
