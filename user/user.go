package user

type User struct {
	*UserConn
	nick string
	user string
}

func CreateUser(nick string, user string, conn *UserConn) *User {
	return &User{conn, nick, user}
}

func (u User) Username() string {
	return u.user
}

func (u User) Nick() string {
	return u.nick
}

func (u *User) SetNick(newNick string) {
	u.nick = newNick
}
