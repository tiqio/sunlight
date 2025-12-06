package core

import "sunlight/cipher"

type TcpConnectionState struct {
	s *Session
}

func (t TcpConnectionState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewTcpConnectionState(s *Session) *TcpConnectionState {
	return &TcpConnectionState{s: s}
}

var _ State = TcpConnectionState{}
