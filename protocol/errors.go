package protocol

type numeric uint16

type IRCError struct {
	num numeric
	msg string
}

var (
	ERR_UNKNOWNCOMMAND  = IRCError{421, "unknown command"}
	ERR_NEEDMOREPARAMS  = IRCError{461, "not enough parameters"}
	ERR_NONICKNAMEGIVEN = IRCError{431, "missing nick"}
)

func (e IRCError) Error() string { return e.msg }

func (e IRCError) Num() numeric { return e.num }
