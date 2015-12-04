// +build windows
package util

import (
	s "syscall"
)

// Initialisation WSA
func initConn() error {
	var d s.WSAData
	return s.WSAStartup(uint32(0x202), &d)
}

func Write(sd s.Handle, sa s.SockaddrInet4, b []byte) error {
	buf := &s.WSABuf{
		Len: uint32(len(b)),
		Buf: &b[0],
	}
	var sent *uint32
	overlapped := s.Overlapped{}
	croutine := byte(0)

	return s.WSASendto(sd, buf, 1, sent, uint32(0), &sa, &overlapped, &croutine)
}
