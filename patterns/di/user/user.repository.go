package user

var users = []User{
	{Id: 123, Username: "Harry"},
	{Id: 456, Username: "Ron"},
	{Id: 789, Username: "Hermione"},
}

type Repository struct{}

func NewUserRepository() Repository {
	return Repository{}
}

func (r Repository) Add(user User) int {
	id := len(users) + 1
	newUser := User{Id: id, Username: user.Username}
	users = append(users, newUser)
	return newUser.Id
}

func (r Repository) GetById(id int) User {
	var user User
	for _, u := range users {
		if u.Id == id {
			user = u
		}
	}
	return user
}

func (r Repository) GetAll() []User {
	return users
}
