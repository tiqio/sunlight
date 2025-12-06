package core

import "sunlight/cipher"

type TerminateState struct {
	s *Session
}

func (t TerminateState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewTerminateState(s *Session) *TerminateState {
	return &TerminateState{s: s}
}

var _ State = TerminateState{}