package core

import "sunlight/cipher"

type UdpRelayState struct {
	s *Session
}

func (u UdpRelayState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpRelayState(s *Session) *UdpRelayState {
	return &UdpRelayState{s: s}
}

var _ State = UdpRelayState{}
