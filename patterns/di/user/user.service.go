package user

type Service struct {
	Repo RepositoryI
}

type ServiceI interface {
	GetAll() []User
	GetById(int) User
	Add(User) int
}

func (s Service) GetAll() []User {
	return s.Repo.GetAll()
}

func (s Service) GetById(id int) User {
	return s.Repo.GetById(id)
}

func (s Service) Add(user User) int {
	return s.Repo.Add(user)
}
