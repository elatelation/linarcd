package protocol

type Nick struct {
	newNick string
}

func IntoNick(ln Line) (Nick, error) {
	if ln.Verb() != "NICK" {
		panic("IntoNick: ln.Verb() != \"NICK\"")
	}
	if len(ln.params) < 1 {
		return Nick{}, ERR_NONICKNAMEGIVEN
	}
	return Nick{ln.params[0]}, nil
}

func (n Nick) Source() *string      { return nil }
func (n Nick) Verb() string         { return "NICK" }
func (n Nick) Parameters() []string { return []string{n.newNick} }
func (n Nick) NewNick() string      { return n.newNick }
