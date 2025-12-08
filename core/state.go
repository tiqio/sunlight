package core

import "net"

type State interface {
	handle(conn net.Conn) error
}

type StateEnum uint8

const (
	INIT StateEnum = iota
	COMMAND
	TCP_CONNECT
	TCP_ACTIVE
	UDP_CONNECT
	UDP_ACTIVE
	TCP_TRANSFER
	UDP_RELAY
	UDP_TRANSFER
	TERMINATE
)
