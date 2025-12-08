package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type TcpConnectionState struct {
	s *Session
}

func (t TcpConnectionState) handle(conn net.Conn) error {
	var nextHop net.Conn
	var err error

	req := t.s.request

	if rule := t.s.getRule(req.Addr.Host); rule == ruleProxy {
		log.Printf(`[socks5] "connect" dial server to connect %s for %s`, req.Addr, conn.RemoteAddr())

		t.s.isProxy = true
		nextHop, err = t.s.dialAccess()
		if err != nil {
			log.Printf(`[socks5] "connect" dial server failed: %s`, err)
			if err = socks.NewReply(socks.HostUnreachable, nil).Write(conn); err != nil {
				log.Printf(`[socks5] "connect" write reply failed: %s`, err)
			}
			return err
		}

	} else {
		log.Printf(`[socks5] "connect" dial %s for %s`, req.Addr, conn.RemoteAddr())

		nextHop, err = net.Dial("tcp", req.Addr.String())
		if err != nil {
			if rule == ruleAuto {
				log.Printf(`[socks5] "connect" dial %s failed, dial server for %s`, req.Addr, conn.RemoteAddr())
				t.s.isProxy = true
				nextHop, err = t.s.dialAccess()
				if err != nil {
					log.Printf(`[socks5] "connect" dial server failed: %s`, err)
					if err = socks.NewReply(socks.HostUnreachable, nil).Write(conn); err != nil {
						log.Printf(`[socks5] "connect" write reply failed: %s`, err)
					}
					return err
				}
				// set proxy in rules
			} else {
				t.s.isProxy = false
				log.Printf(`[socks5] "connect" dial remote failed: %s`, err)
				if err = socks.NewReply(socks.HostUnreachable, nil).Write(conn); err != nil {
					log.Printf(`[socks5] "connect" write reply failed: %s`, err)
				}
				return err
			}
		}
	}

	t.s.nextHop = nextHop
	t.s.setState(TCP_TRANSFER)
	return nil
}

func NewTcpConnectionState(s *Session) *TcpConnectionState {
	return &TcpConnectionState{s: s}
}

var _ State = TcpConnectionState{}
