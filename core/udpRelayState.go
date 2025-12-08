package core

import (
	"net"
)

type UdpRelayState struct {
	s *Session
}

func (u UdpRelayState) handle(conn net.Conn) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpRelayState(s *Session) *UdpRelayState {
	return &UdpRelayState{s: s}
}

var _ State = UdpRelayState{}
