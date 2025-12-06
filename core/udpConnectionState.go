package core

import "sunlight/cipher"

type UdpConnectionState struct {
	s *Session
}

func (u UdpConnectionState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewUdpConnectionState(s *Session) *UdpConnectionState {
	return &UdpConnectionState{s: s}
}

var _ State = &UdpConnectionState{}
