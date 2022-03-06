package data

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type Users []User

var GetUsers = Users{
	{Id: "1", Username: "Harry"},
	{Id: "2", Username: "Hermione"},
	{Id: "3", Username: "Ron"},
}
