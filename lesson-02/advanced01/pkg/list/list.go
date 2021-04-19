package list

type IUser interface {
	Fullname() string
}

type List struct {
	Users []string
}

func New() *List {
	return &List{}
}

func (l *List) AddUser(user IUser) {
	l.Users = append(l.Users, user.Fullname())
}

func (l *List) GetUsers() []string {
	return l.Users
}
