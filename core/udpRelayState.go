package core

import (
	"bytes"
	"io"
	"log"
	"net"
	"sunlight/socks"
	"sunlight/utils"
)

type UdpRelayState struct {
	s *Session
}

func (u UdpRelayState) handle(conn net.Conn) {
	udp := u.s.udp
	nextHop := u.s.nextHop

	addr, _ := socks.NewAddrFromAddr(udp.LocalAddr(), conn.LocalAddr())
	if err := socks.NewReply(socks.Succeeded, addr).Write(conn); err != nil {
		log.Printf(`[socks5] "udp" write reply failed %s`, err)
		u.s.setState(TERMINATE)
	}

	log.Printf(`[socks5] "udp" tunnel established (UDP)%s <-> %s`, udp.LocalAddr(), u.s.accessAddr.String())
	go func() {
		// udp <-> nextHop
		err := relayTunnelUDP(udp, nextHop)
		if err != nil {
			u.s.setState(TERMINATE)
		}
	}()
	if err := waiting4EOF(conn); err != nil {
		u.s.setState(TERMINATE)
		log.Printf(`[socks5] "udp" waiting for EOF failed: %s`, err)
	}
	log.Printf(`[socks5] "udp" tunnel disconnected (UDP)%s >-< %s`, udp.LocalAddr(), u.s.accessAddr.String())
}

func relayTunnelUDP(udp net.PacketConn, conn net.Conn) error {
	errc := make(chan error, 2)
	var clientAddr net.Addr

	go func() {
		b := utils.LPool.Get().([]byte)
		defer utils.LPool.Put(b)

		for {
			n, addr, err := udp.ReadFrom(b)
			if err != nil {
				errc <- err
				return
			}

			dgram, err := socks.ReadUDPDatagram(bytes.NewReader(b[:n]))
			if err != nil {
				errc <- err
				return
			}
			if clientAddr == nil {
				clientAddr = addr
			}
			dgram.Header.Rsv = uint16(len(dgram.Data))
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

			if clientAddr == nil {
				continue
			}
			dgram.Header.Rsv = 0
			buf := bytes.NewBuffer(nil)
			dgram.Write(buf)
			if _, err := udp.WriteTo(buf.Bytes(), clientAddr); err != nil {
				errc <- err
				return
			}
		}
	}()

	return <-errc
}

func waiting4EOF(conn net.Conn) (err error) {
	b := utils.SPool.Get().([]byte)
	defer utils.SPool.Put(b)
	for {
		_, err = conn.Read(b)
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
	}
	return
}

func NewUdpRelayState(s *Session) *UdpRelayState {
	return &UdpRelayState{s: s}
}

var _ State = UdpRelayState{}
