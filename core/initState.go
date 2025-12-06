package core

import "sunlight/cipher"

type InitState struct {
	s *Session
}

func (i InitState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewInitState(s *Session) *InitState {
	return &InitState{s: s}
}

var _ State = InitState{}
