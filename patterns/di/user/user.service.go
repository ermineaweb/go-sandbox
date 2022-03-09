package user

type Service struct {
	Repo RepositoryI
}

type RepositoryI interface {
	Add(User) int
	GetById(int) User
	GetAll() []User
}

func NewUserService(repository RepositoryI) Service {
	return Service{Repo: repository}
}

func (s Service) Add(user User) int {
	return s.Repo.Add(user)
}

func (s Service) GetById(id int) User {
	return s.Repo.GetById(id)
}

func (s Service) GetAll() []User {
	return s.Repo.GetAll()
}
