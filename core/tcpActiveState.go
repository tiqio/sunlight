package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type TcpActiveState struct {
	s *Session
}

func (t TcpActiveState) handle(conn net.Conn) {
	req := t.s.request
	log.Printf(`[socks5] "connect" connect %s for %s`, req.Addr, conn.RemoteAddr())
	nextHop, err := net.Dial("tcp", req.Addr.String())
	if err != nil {
		log.Printf(`[socks5] "connect" dial remote failed: %s`, err)
		if err := socks.NewReply(socks.HostUnreachable, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "connect" write reply failed: %s`, err)
		}
		t.s.setState(TERMINATE)
	}

	t.s.isProxy = false
	t.s.nextHop = nextHop
	t.s.setState(TCP_TRANSFER)
}

func NewTcpActiveState(s *Session) *TcpActiveState {
	return &TcpActiveState{s: s}
}

var _ State = TcpActiveState{}
