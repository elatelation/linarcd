package protocol

type User struct {
	user string
	real string
}

func IntoUser(ln Line) (User, error) {
	if ln.cmd != "USER" {
		panic("IntoUser: ln.cmd != \"USER\"")
	}
	if len(ln.params) < 4 {
		return User{}, ERR_NEEDMOREPARAMS
	}
	user := ln.params[0]
	real := ln.params[3]
	return User{user, real}, nil
}

func (u User) Source() *string      { return nil }
func (u User) Verb() string         { return "USER" }
func (u User) Parameters() []string { return []string{u.user, "0", "*", u.real} }
func (u User) Name() string         { return u.user }
func (u User) Real() string         { return u.real }
