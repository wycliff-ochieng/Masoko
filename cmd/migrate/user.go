package migrate

type User struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
}

var Users = []*User{
	&User{}, &User{}, &User{}, &User{},
}

func GetUsers() []*User {
	return Users
}

func CreateUser(u *User) {
	u.ID = getNextUser()
	Users = append(Users, u)
}

func getNextUser() int {
	lastUser := Users[len(Users)-1]
	return lastUser.ID + 1
}
