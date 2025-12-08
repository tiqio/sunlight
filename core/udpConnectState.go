package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type UdpConnectionState struct {
	s *Session
}

func (u UdpConnectionState) handle(conn net.Conn) {
	log.Printf(`[socks5] "udp" associate UDP for %s`, conn.RemoteAddr())
	udp, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Printf(`[socks5] "udp" UDP associate failed on listen: %s`, err)
		if err := socks.NewReply(socks.Failure, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "udp" write reply failed %s`, err)
		}
		u.s.setState(TERMINATE)
	}

	nextHop, err := u.s.requestServer4UDP()
	if err != nil {
		log.Printf(`[socks5] "udp" UDP associate failed on request the server: %s`, err)
		if err := socks.NewReply(socks.Failure, nil).Write(conn); err != nil {
			log.Printf(`[socks5] "udp" Write reply failed %s`, err)
		}

		u.s.setState(TERMINATE)
	}

	u.s.udp = udp
	u.s.nextHop = nextHop
	u.s.setState(UDP_RELAY)
}

func NewUdpConnectionState(s *Session) *UdpConnectionState {
	return &UdpConnectionState{s: s}
}

var _ State = &UdpConnectionState{}
