package main

import (
	"fmt"
	"gobox/server"
	"io/ioutil"
	"os"
	s "syscall"
)

const (
	PORT = 3000
)

func main() {
	var sa s.SockaddrInet4
	fmt.Println(sa)
	fd, err := s.Socket(s.AF_INET, s.SOCK_STREAM, 0)
	if err != nil {
		check(err)
	}
	defer s.Close(fd)
	if err := s.Bind(fd, &s.SockaddrInet4{Port: PORT, Addr: [4]byte{0, 0, 0, 0}}); err != nil {
		check(err)
	}

	if err := s.Listen(fd, 5); err != nil {
		check(err)
	}

	for {
		nfd, sa, err := s.Accept(fd)
		if err != nil {
			check(err)
		}

		go func(nfd int, sa s.Sockaddr) {
			fmt.Println(sa)
			defer s.Close(nfd)
			b := make([]byte, 500)
			var n int
			for {
				n, err = s.Read(nfd, b)
				if n != 0 {
					break
				}

			}
			fmt.Println(n)
			err := ioutil.WriteFile("./test.txt", b, 0644)

			fmt.Println("Fichier créé")
			check(err)

		}(nfd, sa)
	}
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
