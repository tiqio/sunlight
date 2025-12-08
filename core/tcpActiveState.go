package core

import (
	"net"
)

type TcpActiveState struct {
	s *Session
}

func (t TcpActiveState) handle(conn net.Conn) error {
	//TODO implement me
	panic("implement me")
}

func NewTcpActiveState(s *Session) *TcpActiveState {
	return &TcpActiveState{s: s}
}

var _ State = TcpActiveState{}
