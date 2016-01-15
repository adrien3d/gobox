// POSIX n'existant pas sous windows, voici l'équivalent avec les librairys WSA.
package util

import (
	s "syscall"
)

// Structure d'un dossier
type Conn struct {
	sd s.Handle
	sa s.SockaddrInet4
}

// Etablie une connexion via socket à l'aide des librairys WSA (windows).
func (c *Conn) Dial(port int, addr [4]byte) (err error) {
	c.sa = s.SockaddrInet4{Port: port, Addr: addr}
	var d s.WSAData
	err = s.WSAStartup(uint32(0x202), &d)
	if err != nil {
		return
	}
	c.sd, err = s.Socket(s.AF_INET, s.SOCK_STREAM, 0)
	if err != nil {
		return
	}
	s.Connect(c.sd, &c.sa)
	return
}

func (c *Conn) Write(b []byte) error {
	buf := &s.WSABuf{
		Len: uint32(len(b)),
		Buf: &b[0],
	}
	var sent *uint32
	overlapped := s.Overlapped{}
	croutine := byte(0)

	return s.WSASendto(c.sd, buf, 1, sent, uint32(0), &c.sa, &overlapped, &croutine)
}

func (c *Conn) Read(b []byte) (int, error) {
	dataBuf := s.WSABuf{Len: uint32(len(b)), Buf: &b[0]}
	flags := uint32(0)
	qty := uint32(0)
	//fmt.Println(c.sd, &dataBuf, 1, &qty, &flags)
	err := s.WSARecv(c.sd, &dataBuf, 1, &qty, &flags, nil, nil)
	return int(qty), err
}

func (c *Conn) Close() {
	s.Close(c.sd)
}
