package core

import "sunlight/cipher"

type State interface {
	handleFrame(frame cipher.Frame) error
}
