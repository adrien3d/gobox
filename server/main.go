package main

import (
	"encoding/binary"
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
	s "syscall"
)

const (
	PORT = 1002
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
			fmt.Println("Nouveau socket :\n\t", sa)

			defer s.Close(nfd)
			size := make([]byte, 8)
			_, err := s.Read(nfd, size)
			check(err)
			lenght := binary.BigEndian.Uint64(size) // voir le code sur le serveur, retourne la valeur max de int64 (ou presque)
			file := make([]byte, lenght)
			f, err := os.Create("./config.json")
			check(err)
			defer f.Close()
			for {
				fmt.Println(lenght, " octets restants.")
				if lenght-util.MAXSIZE > 0 {
					b := make([]byte, util.MAXSIZE)
					_, err := s.Read(nfd, b)
					check(err)
					file = append(file, b...)
					lenght = lenght - util.MAXSIZE
				} else {
					b := make([]byte, lenght)
					_, err := s.Read(nfd, b)
					check(err)
					file = append(file, b...)
					break
				}

			}

			err = ioutil.WriteFile("./config.json", file, 0644)
			check(err)
			clientListRep, err := util.BytesToFol(file)
			check(err)

			fmt.Println("Arborescence client :")
			fmt.Println(clientListRep.ToString())

			//err := ioutil.WriteFile("./test.txt", b, 0644)

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
