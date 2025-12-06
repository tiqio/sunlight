package core

import (
	"sunlight/cipher"
)

type UdpActiveState struct {
	s *Session
}

func (u UdpActiveState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpActiveState(s *Session) *UdpActiveState {
	return &UdpActiveState{s: s}
}

var _ State = UdpActiveState{}
