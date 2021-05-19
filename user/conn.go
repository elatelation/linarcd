package user

import (
	"bufio"
	"errors"
	"github.com/elatelation/linarcd/protocol"
	"net"
	"os"
	"time"
)

type PingError struct{}

var (
	PingTimeout = PingError{}
)

func (PingError) Error() string {
	return "ping timeout (2 minutes)"
}

type UserConn struct {
	net.Conn
	rd              *bufio.Reader
	pingedNotPonged bool
}

func NewUserConn(s net.Conn) *UserConn {
	rd := bufio.NewReaderSize(s, 512)
	return &UserConn{s, rd, false}
}

// func (uc *UserConn) ReadLine() (protocol.Message, error) {
// 	s, err := uc.rd.ReadString('\n')
// 	if err != nil {
// 		return nil, err
// 	}
// 	return protocol.Parse(s)
// }

func (uc *UserConn) ReadString() (string, error) {
	if err := uc.SetReadDeadline(time.Now().Add(time.Minute)); err != nil {
		return "", err
	}
	s, err := uc.rd.ReadString('\n')
	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			if uc.pingedNotPonged {
				return "", PingTimeout
			}
			if _, err = uc.Write([]byte("PING\r\n")); err != nil {
				return "", err
			}
			uc.pingedNotPonged = true
			return uc.ReadString()
		} else {
			return "", err
		}
	}
	uc.pingedNotPonged = false
	return s, nil
}

func (uc *UserConn) Close() {
	uc.rd = nil
	if err := uc.Conn.Close(); err != nil {
		panic(err.Error())
	}
}

func (uc *UserConn) Send(msg protocol.Message) error {
	_, err := uc.Write(protocol.ToBytes(msg, false))
	return err
}
