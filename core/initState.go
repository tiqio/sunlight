package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type InitState struct {
	s *Session
}

func (i InitState) handle(conn net.Conn) error {
	// select method
	methods, err := socks.ReadMethods(conn)
	if err != nil {
		log.Printf(`[socks5] read methods failed: %s`, err)
		return err
	}
	method := socks.MethodNoAcceptable
	for _, m := range methods {
		if m == socks.MethodNoAuth {
			method = m
		}
	}

	if err := socks.WriteMethod(method, conn); err != nil || method == socks.MethodNoAcceptable {
		if err != nil {
			log.Printf(`[socks5] write method failed: %s`, err)
		} else {
			log.Printf(`[socks5] methods is not acceptable`)
		}
		return err
	}

	i.s.setState(COMMAND)
	return nil
}

func NewInitState(s *Session) *InitState {
	return &InitState{s: s}
}

var _ State = InitState{}
