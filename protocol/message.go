package protocol

import (
	"fmt"
	"strings"
)

type Message interface {
	Source() *string
	Verb() string
	Parameters() []string
}

func Parse(s string) (Message, error) {
	line, err := ParseLine(s)
	if err != nil {
		return nil, err
	}
	switch line.Verb() {
	case "NICK":
		return IntoNick(line)
	case "USER":
		return IntoUser(line)
	default:
		return line, nil
	}
}

func ToBytes(msg Message, trail bool) []byte {
	var builder strings.Builder
	var trailer string
	if src := msg.Source(); src != nil {
		builder.WriteByte(':')
		builder.WriteString(*src)
		builder.WriteByte(' ')
	}
	fmt.Println(builder.String())
	builder.WriteString(msg.Verb())
	fmt.Println(builder.String())
	params := msg.Parameters()
	if len(params) > 0 {
		trailer, params = params[len(params)-1], params[:len(params)-1]
	}
	for _, param := range params {
		fmt.Println(param)
		builder.WriteByte(' ')
		builder.WriteString(param)
		fmt.Println(builder.String())
	}
	if trailer != "" {
		if trail || strings.IndexByte(trailer, ' ') != -1 {
			builder.Write([]byte(" :"))
			builder.WriteString(trailer)
		} else {
			builder.WriteByte(' ')
			builder.WriteString(trailer)
		}
	}
	fmt.Println(builder.String())
	b := []byte(builder.String())
	if len(b) >= 512 {
		b = b[:512-2]
	}
	b = append(b, "\r\n"...)
	return b
}
