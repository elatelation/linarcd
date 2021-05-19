package main

import (
	"fmt"
	"github.com/elatelation/linarcd/protocol"
	"github.com/elatelation/linarcd/user"
	"io"
	"net"
	"os"
	"strings"
)

var serverName string = "unknown"

func handleNewConnection(c net.Conn, ch chan<- *user.User) {
	var (
		nick     *string
		username *string
	)
	uc := user.NewUserConn(c)
	for {
		s, err := uc.ReadString()
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Printf("%s disconnected before registration\n", c.RemoteAddr().String())
			default:
				fmt.Printf("read: %s\n", err.Error())
			}
			uc.Close()
			return
		}
		ln, err := protocol.Parse(s)
		if err != nil && err == protocol.ParseErr {
			continue
		}
		fmt.Printf("%#v\n", ln)
		if src := ln.Source(); src != nil {
			fmt.Printf("src = %v\n", *src)
		}
		if nickMsg, ok := ln.(protocol.Nick); ok {
			_nick := nickMsg.NewNick()
			nick = &_nick

		}
		if userMsg, ok := ln.(protocol.User); ok {
			_user := userMsg.Name()
			username = &_user
		}
		if nick != nil && username != nil {
			uc.Send(protocol.CreateWelcome(serverName, *nick))
			uc.Send(protocol.CreateYourHost(serverName, *nick))
			uc.Send(protocol.CreateCreated(serverName, *nick))
			uc.Send(protocol.CreateMyInfo(serverName, *nick))
			u := user.CreateUser(*nick, *username, uc)
			ch <- u
			go mainUserLoop(u)
			return
		}
	}
}

func mainUserLoop(u *user.User) {
	for {
		s, err := u.ReadString()
		if err != nil {
			switch err {
			case io.EOF:
				fmt.Printf("%s!%s@%s reset\n", u.Nick(), u.Username(), u.RemoteAddr().String())
			default:
				fmt.Printf("recv: %s\n", err.Error())
			}
			u.Close()
			return
		}
		msg, err := protocol.Parse(s)
		if err != nil && err == protocol.ParseErr {
		}
		_ = msg
		panic("unimplemented")

	}
}

func createUserMonitor() chan<- *user.User {
	ch := make(chan *user.User)
	go func() {
		users := make(map[string]*user.User) // usernames dont change
		nicks := make(map[string]string)     // nick -> user
		for u := range ch {
			users[u.Username()] = u
			nicks[strings.ToLower(u.Nick())] = u.Username()
		}
	}()
	return ch
}

func listen() {
	listener, err := net.Listen("tcp", ":6667")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("listening on %s\n", listener.Addr().String())
	userCh := createUserMonitor()
	failedOnce := false
	for {
		s, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept: %s\n", err.Error())
			if failedOnce {
				os.Exit(1)
			} else {
				failedOnce = true
			}
		} else {
			fmt.Printf("accepted from %s\n", s.RemoteAddr().String())
			go handleNewConnection(s, userCh)
		}
	}
}

func main() {
	if len(os.Args) > 1 {
		line := os.Args[1]
		msg, err := protocol.Parse(line)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%#v\n", msg)
		if src := msg.Source(); src != nil {
			fmt.Printf("Source() = %s\n", *src)
		}
		fmt.Printf("%v\n", protocol.ToBytes(msg, false))
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}
	serverName = hostname
	listen()
}
