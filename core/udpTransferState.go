package core

import "sunlight/cipher"

type UdpTransferState struct {
	s *Session
}

func (u UdpTransferState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpTransferState(s *Session) *UdpTransferState {
	return &UdpTransferState{s: s}
}

var _ State = UdpTransferState{}
