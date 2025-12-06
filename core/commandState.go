package core

import (
	"sunlight/cipher"
)

type CommandState struct {
	s *Session
}

func (c CommandState) handleFrame(frame cipher.Frame) error {
	//TODO implement me
	panic("implement me")
}

func NewCommandState(s *Session) *CommandState {
	return &CommandState{s: s}
}

var _ State = CommandState{}
