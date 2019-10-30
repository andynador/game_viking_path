package models

type User struct {
	id       int
	login    string
	warriors []*Warrior
}

func NewUser(id int, login string) *User {
	return &User{
		id:       id,
		login:    login,
		warriors: make([]*Warrior, 0),
	}
}

func (user *User) GetID() int {
	return user.id
}

func (user *User) GeLogin() string {
	return user.login
}

func (user *User) AddWarrior(warrior *Warrior) *User {
	user.warriors = append(user.warriors, warrior)

	return user
}

func (user *User) GetWarriors() []*Warrior {
	return user.warriors
}
