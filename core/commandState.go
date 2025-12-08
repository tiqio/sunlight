package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type CommandState struct {
	s *Session
}

func (c CommandState) handle(conn net.Conn) {
	// read command
	request, err := socks.ReadRequest(conn)
	if err != nil {
		log.Printf(`[socks5] read command failed: %s`, err)
		c.s.setState(TERMINATE)
	}
	c.s.request = request
	switch request.Cmd {
	case socks.CmdConnect:
		c.s.setState(TCP_CONNECT)
	case socks.CmdUDP:
		c.s.setState(UDP_CONNECT)
	default:
		log.Fatalf("[socks5] unknown command: %d", request.Cmd)
	}
	return
}

func NewCommandState(s *Session) *CommandState {
	return &CommandState{s: s}
}

var _ State = CommandState{}
