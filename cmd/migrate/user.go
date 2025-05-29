package migrate

import "github.com/wycliff-ochieng/db"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
}

type RegisterUserPayload struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
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

func (p *db.Postgrestore) GetUserByEmail(email string) (*User, error) {
	row, err := p.db.Query()
}

func ScanRowIntoUsers()
