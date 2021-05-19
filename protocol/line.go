package protocol

import "strings"

type ParseError struct{}

func (ParseError) Error() string { return "not an irc line" }

var (
	ParseErr = ParseError{}
)

// we dont do tags
type Line struct {
	src    *string
	cmd    string
	params []string
}

func ParseLine(s string) (Line, error) {
	var (
		src    *string
		cmd    string
		params []string
	)
	tokens := strings.Split(strings.TrimRight(s, "\r\n"), " ")
	if len(tokens) < 1 {
		return Line{}, ParseErr
	}
	token, tokens := tokens[0], tokens[1:]
	if len(token) < 1 {
		return Line{}, ParseErr
	}
	if token[0] == ':' {
		token = token[1:]
		src = &token
		if len(tokens) < 1 {
			return Line{}, ParseErr
		}
		cmd, tokens = tokens[0], tokens[1:]
	} else {
		cmd = token
	}
	for i, param := range tokens {
		if param[0] == ':' {
			trailing := strings.Join(tokens[i:], " ")
			params = append(params, trailing[1:])
			break
		}
		params = append(params, param)
	}
	return Line{src, cmd, params}, nil
}

func NewLine(src string, cmd string, params []string) Line {
	if src != "" {
		return Line{&src, cmd, params}
	} else {
		return Line{nil, cmd, params}
	}
}

func (r Line) Source() *string      { return r.src }
func (r Line) Verb() string         { return r.cmd }
func (r Line) Parameters() []string { return r.params }
