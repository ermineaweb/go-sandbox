package user

// public interface
type User interface {
	GetId() int
	SetId(int)
	GetUsername() string
	SetUsername(string)
}

// private object
type user struct {
	id       int
	username string
}

// public method to construct the object
func NewUser(username string) User {
	return user{username: username}
}

// publics getters / setters
func (u user) GetId() int {
	return u.id
}

func (u user) SetId(id int) {
	u.id = id
}

func (u user) GetUsername() string {
	return u.username
}

func (u user) SetUsername(username string) {
	u.username = username
}
