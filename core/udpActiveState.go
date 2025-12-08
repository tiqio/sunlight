package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type UdpActiveState struct {
	s *Session
}

func (u UdpActiveState) handle(conn net.Conn) {
	log.Printf(`[socks5] "udp-over-tcp" associate UDP for %s`, conn.RemoteAddr())
	udp, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Printf(`[socks5] "udp-over-tcp" UDP associate failed on listen: %s`, err)
		if err := socks.NewReply(socks.Failure, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "udp-over-tcp" write reply failed %s`, err)
		}
		u.s.setState(TERMINATE)
	}

	u.s.udp = udp
	u.s.setState(UDP_TRANSFER)
}

func NewUdpActiveState(s *Session) *UdpActiveState {
	return &UdpActiveState{s: s}
}

var _ State = UdpActiveState{}
