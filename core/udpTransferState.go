package core

import (
	"log"
	"net"
	"sunlight/socks"
	"sunlight/utils"
)

type UdpTransferState struct {
	s *Session
}

func (u UdpTransferState) handle(conn net.Conn) {
	udp := u.s.udp
	addr, _ := socks.NewAddrFromAddr(udp.LocalAddr(), conn.LocalAddr())
	if err := socks.NewReply(socks.Succeeded, addr).Write(conn); err != nil {
		log.Printf(`[socks5] "udp-over-tcp" write reply failed %s`, err)
		u.s.setState(TERMINATE)
	}

	log.Printf(`[socks5] "udp-over-tcp" tunnel established %s <-> (UDP)%s`, conn.RemoteAddr(), udp.LocalAddr())
	if err := transferTunnelUDP(conn, udp); err != nil {
		log.Printf(`[socks5] "udp-over-tcp" tunnel UDP failed: %s`, err)
	}
	log.Printf(`[socks5] "udp-over-tcp" tunnel disconnected %s >-< (UDP)%s`, conn.RemoteAddr(), udp.LocalAddr())
}

func transferTunnelUDP(conn net.Conn, udp net.PacketConn) error {
	errc := make(chan error, 2)

	go func() {
		b := utils.LPool.Get().([]byte)
		defer utils.LPool.Put(b)

		for {
			n, addr, err := udp.ReadFrom(b)
			if err != nil {
				errc <- err
				return
			}

			saddr, _ := socks.NewAddr(addr.String())
			dgram := socks.NewUDPDatagram(
				socks.NewUDPHeader(uint16(n), 0, saddr), b[:n])
			if err := dgram.Write(conn); err != nil {
				errc <- err
				return
			}
		}
	}()

	go func() {
		for {
			dgram, err := socks.ReadUDPDatagram(conn)
			if err != nil {
				errc <- err
				return
			}

			addr, err := net.ResolveUDPAddr("udp", dgram.Header.Addr.String())
			if err != nil {
				continue
			}
			if _, err := udp.WriteTo(dgram.Data, addr); err != nil {
				errc <- err
				return
			}
		}
	}()

	return <-errc
}

func NewUdpTransferState(s *Session) *UdpTransferState {
	return &UdpTransferState{s: s}
}

var _ State = UdpTransferState{}
