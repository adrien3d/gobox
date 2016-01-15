// POSIX n'existant pas sous windows, voici l'équivalent avec les librairys WSA.
package util

import (
	"fmt"
	"sync"
	s "syscall"
)

// Structure d'un dossier
type Conn struct {
	sd         s.Handle
	sa         s.SockaddrInet4
	wsadata    s.WSAData
	overlapped s.Overlapped
	sync       *sync.Mutex
}

// Etablie une connexion via socket à l'aide des librairys WSA (windows).
func (c *Conn) Dial(port int, addr [4]byte) (err error) {
	c.sa = s.SockaddrInet4{Port: port, Addr: addr}
	fmt.Println("OVERLAPPED")
	fmt.Println(c.overlapped)
	err = s.WSAStartup(uint32(0x202), &c.wsadata)
	if err != nil {
		return
	}
	c.sd, err = s.Socket(s.AF_INET, s.SOCK_STREAM, 6)
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
	c.overlapped = s.Overlapped{}
	croutine := byte(0)
	fmt.Printf("\nEnvoi en cours\n")
	err := s.WSASendto(c.sd, buf, 1, sent, uint32(0), &c.sa, &c.overlapped, &croutine)
	return err
}

func (c *Conn) Read(b []byte) (int, error) {
	lenght := uint32(len(b))
	dataBuf := &s.WSABuf{Len: lenght, Buf: &b[0]}
	flags := uint32(0)
	qty := uint32(0)

	fmt.Printf("\nReception en cours\n")
	err := s.WSARecv(c.sd, dataBuf, 1, &qty, &flags, nil, nil)
	return int(qty), err
}

func (c *Conn) Close() {
	s.WSACleanup()
	//s.Closesocket(c.sa)
	s.Close(c.sd)
}

func (c *Conn) Lock() {
	c.sync.Lock()
}
