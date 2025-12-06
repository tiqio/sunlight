package core

import "sunlight/cipher"

type TcpActiveState struct {
	s *Session
}

func (t TcpActiveState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewTcpActiveState(s *Session) *TcpActiveState {
	return &TcpActiveState{s: s}
}

var _ State = TcpActiveState{}
