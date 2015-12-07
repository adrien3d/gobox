package util

import (
	"fmt"
	s "syscall"
)

// Structure d'un dossier
type Conn struct {
	sd s.Handle
	sa s.SockaddrInet4
}

// Etablie une connexion via socket à l'aide des librairies WSA (windows).
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

	fmt.Println("PLOP")
	return s.WSASendto(c.sd, buf, 1, sent, uint32(0), &c.sa, &overlapped, &croutine)
}

func (c *Conn) Close() {
	s.Close(c.sd)
}
