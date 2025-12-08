package core

import (
	"net"
)

type UdpActiveState struct {
	s *Session
}

func (u UdpActiveState) handle(conn net.Conn) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpActiveState(s *Session) *UdpActiveState {
	return &UdpActiveState{s: s}
}

var _ State = UdpActiveState{}
