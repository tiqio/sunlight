package core

import "sunlight/cipher"

type TcpTransferState struct {
	s *Session
}

func (t TcpTransferState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewTcpTransferState(s *Session) *TcpTransferState {
	return &TcpTransferState{s: s}
}

var _ State = TcpTransferState{}