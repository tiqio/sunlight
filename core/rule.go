package core

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sunlight/socks"
)

const (
	ruleNone = iota
	ruleProxy
	ruleDirect
	ruleAuto
)

func (s *Session) getRule(addr string) int {
	return ruleProxy
}

func (s *Session) dialAccess() (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", s.accessAddr.IP, s.accessAddr.Port))
	if err != nil {
		return nil, err
	}

	// handshake
	if err := socks.WriteMethods([]byte{socks.MethodNoAuth}, conn); err != nil {
		conn.Close()
		return nil, err
	}
	buf := make([]byte, 2)
	if _, err := io.ReadFull(conn, buf); err != nil {
		conn.Close()
		return nil, err
	}
	if buf[0] != socks.Version || buf[1] != socks.MethodNoAuth {
		conn.Close()
		return nil, errors.New("Handshake failed")
	}

	return conn, nil
}

func (s *Session) requestServer4UDP() (net.Conn, error) {
	ser, err := s.dialAccess()
	if err != nil {
		return nil, err
	}

	if err := socks.NewRequest(socks.CmdUDPOverTCP, nil).Write(ser); err != nil {
		ser.Close()
		return nil, err
	}
	res, err := socks.ReadReply(ser)
	if err != nil {
		ser.Close()
		return nil, err
	}
	if res.Rep != socks.Succeeded {
		ser.Close()
		return nil, fmt.Errorf("Request UDP over TCP associate failed: %q", res.Rep)
	}
	return ser, nil
}
