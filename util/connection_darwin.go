// Ce package utilise les fonctions POSIX pour communiquer via socket.
// Il a été recréé dans un but pédagogique, mais c'est "l'équivalent"
// du package "net", en plus sommaire.
package util

import (
	s "syscall"
)

// Structure d'un dossier
type Conn struct {
	sd int
	sa s.SockaddrInet4
}

// Etablie une connexion via socket à l'aide des librairies POSIX (BSD)
func (c *Conn) Dial(port int, addr [4]byte) (err error) {
	c.sa = s.SockaddrInet4{Port: port, Addr: addr}
	c.sd, err = s.Socket(s.AF_INET, s.SOCK_STREAM, 0)
	if err != nil {
		return
	}
	s.Connect(c.sd, &c.sa)
	return
}

func (c *Conn) Write(b []byte) error {
	_, err := s.Write(c.sd, b)
	return err
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return s.Read(c.sd, b)
}

func (c *Conn) Close() {
	s.Close(c.sd)
}
