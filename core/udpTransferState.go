package core

import (
	"net"
)

type UdpTransferState struct {
	s *Session
}

func (u UdpTransferState) handle(conn net.Conn) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpTransferState(s *Session) *UdpTransferState {
	return &UdpTransferState{s: s}
}

var _ State = UdpTransferState{}
