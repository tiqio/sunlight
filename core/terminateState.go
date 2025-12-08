package core

import (
	"net"
)

type TerminateState struct {
	s *Session
}

func (t TerminateState) handle(conn net.Conn) {
	if t.s.nextHop != nil {
		if _, ok := t.s.nextHop.(net.Conn); ok {
			t.s.nextHop.Close()
		}
		t.s.nextHop = nil
	}

	if t.s.udp != nil {
		t.s.udp.Close()
		t.s.udp = nil
	}

	if conn != nil {
		conn.Close()
	}

	t.s.setState(INIT)
	t.s.request = nil
}

func NewTerminateState(s *Session) *TerminateState {
	return &TerminateState{s: s}
}

var _ State = TerminateState{}
