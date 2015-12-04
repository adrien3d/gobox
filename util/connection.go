package util

import (
	s "syscall"
)

const (
	MAXSIZE = 500
	PORT    = 3000
)

var (
	ADDR = [4]byte{10, 8, 0, 1}
)

// Etablie une connexion via socket sur le serveur
// à l'aide des librairies POSIX ou WSA.
func Dial() (s.Handle, s.SockaddrInet4, error) {
	var sa s.SockaddrInet4 = s.SockaddrInet4{Port: PORT, Addr: ADDR}
	var sd s.Handle
	err := initConn()
	if err != nil {
		return sd, sa, err
	}
	sd, err = s.Socket(s.AF_INET, s.SOCK_STREAM, 0)
	if err != nil {
		return sd, sa, err
	}
	s.Connect(sd, &sa)
	return sd, sa, err
}
