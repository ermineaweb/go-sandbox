package user

var mockUsers = []User{
	{Id: 44, Username: "Homer"},
	{Id: 55, Username: "Bart"},
	{Id: 66, Username: "Lisa"},
}

type MockRepository struct{}

func NewMockRepository() MockRepository {
	return MockRepository{}
}

func (r MockRepository) Add(user User) int {
	id := len(mockUsers) + 1
	newUser := User{Id: id, Username: user.Username}
	mockUsers = append(mockUsers, newUser)
	return newUser.Id
}

func (r MockRepository) GetById(id int) User {
	var user User
	for _, u := range mockUsers {
		if u.Id == id {
			user = u
		}
	}
	return user
}

func (r MockRepository) GetAll() []User {
	return mockUsers
}
