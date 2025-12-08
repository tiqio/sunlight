package core

import (
	"net"
)

type TerminateState struct {
	s *Session
}

func (t TerminateState) handle(conn net.Conn) error {
	//TODO implement me
	panic("implement me")
}

func NewTerminateState(s *Session) *TerminateState {
	return &TerminateState{s: s}
}

var _ State = TerminateState{}
