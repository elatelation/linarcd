package protocol

import (
	"fmt"
)

const (
	RPL_WELCOME  = 1
	RPL_YOURHOST = 2
	RPL_CREATED  = 3
	RPL_MYINFO   = 4
)

func CreateWelcome(server string, nick string) Line {
	return NewLine(server, fmt.Sprintf("%03d", RPL_WELCOME), []string{nick, fmt.Sprintf("welcome to irc %s", nick)})
}

func CreateYourHost(server string, nick string) Line {
	return NewLine(
		server,
		fmt.Sprintf("%03d", RPL_YOURHOST),
		[]string{
			nick,
			fmt.Sprintf("you are connected to %s running linarcd", server),
		},
	)
}

func CreateCreated(server, nick string) Line {
	return NewLine(
		server,
		fmt.Sprintf("%03d", RPL_CREATED),
		[]string{
			nick,
			"created who knows when",
		},
	)
}

func CreateMyInfo(server, nick string) Line {
	return NewLine(
		server,
		fmt.Sprintf("%03d", RPL_MYINFO),
		[]string{
			nick,
			server,
			"0",
			"",
			"",
			"",
		},
	)
}
