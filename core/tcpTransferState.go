package core

import (
	"log"
	"net"
	"sunlight/socks"
	"sunlight/utils"
)

type TcpTransferState struct {
	s *Session
}

func (t TcpTransferState) handle(conn net.Conn) error {
	var dash rune
	req := t.s.request
	nextHop := t.s.nextHop

	if t.s.isProxy {
		if err := req.Write(nextHop); err != nil {
			log.Printf(`[socks5] "connect" send request failed: %s`, err)
			return err
		}
		dash = '-'
	} else {
		if err := socks.NewReply(socks.Succeeded, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "connect" write reply failed: %s`, err)
			return err
		}
		dash = '='
	}

	log.Printf(`[socks5] "connect" tunnel established %s <%c> %s`, conn.RemoteAddr(), dash, req.Addr)
	if err := utils.Transport(conn, nextHop); err != nil {
		log.Printf(`[socks5] "connect" transport failed: %s`, err)
	}
	log.Printf(`[socks5] "connect" tunnel disconnected %s >%c< %s`, conn.RemoteAddr(), dash, req.Addr)
	return nil
}

func NewTcpTransferState(s *Session) *TcpTransferState {
	return &TcpTransferState{s: s}
}

var _ State = TcpTransferState{}
