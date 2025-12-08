package core

import (
	"log"
	"net"
	"sunlight/socks"
)

type Session struct {
	id int

	currentState State
	stateMap     map[StateEnum]State

	request *socks.Request
	isProxy bool
	nextHop net.Conn
	udp     net.Conn

	accessAddr net.TCPAddr
}

func (s *Session) setState(enum StateEnum) {
	state, exists := s.stateMap[enum]
	if !exists {
		log.Fatalf("Session setState(%d): unknown state", enum)
		return
	}
	s.currentState = state
}

func (s *Session) registerState(enum StateEnum, state State) {
	s.stateMap[enum] = state
}

func NewSession(id int) *Session {
	s := &Session{id: id, stateMap: make(map[StateEnum]State)}

	initState := NewInitState(s)
	commandState := NewCommandState(s)
	tcpConnectState := NewTcpConnectionState(s)
	tcpActiveState := NewTcpActiveState(s)
	udpConnectState := NewUdpConnectionState(s)
	udpActiveState := NewUdpActiveState(s)
	tcpTransferState := NewTcpTransferState(s)
	udpRelayState := NewUdpRelayState(s)
	udpTransferState := NewUdpTransferState(s)

	s.registerState(INIT, initState)
	s.registerState(COMMAND, commandState)
	s.registerState(TCP_CONNECT, tcpConnectState)
	s.registerState(TCP_ACTIVE, tcpActiveState)
	s.registerState(UDP_CONNECT, udpConnectState)
	s.registerState(UDP_ACTIVE, udpActiveState)
	s.registerState(TCP_TRANSFER, tcpTransferState)
	s.registerState(UDP_RELAY, udpRelayState)
	s.registerState(UDP_TRANSFER, udpTransferState)

	return s
}
