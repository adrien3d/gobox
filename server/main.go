package main

import (
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

			// On récupère le nombre d'octets qui va être reçu
			size := make([]byte, 8)
			_, err := s.Read(nfd, size)
			check(err)
			var lenght int64 = bigInt(size, 0) // conversion des 8 octets en entier

			// Création du buffer de réception
			file := make([]byte, 0, lenght+1)
			check(err)

			// Réception et assemblage
			for {
				fmt.Println(lenght, " octets restants.")
				if lenght == 0 {
					break
				}

				b := make([]byte, util.MAXSIZE)
				n, err := s.Read(nfd, b)
				check(err)
				fmt.Println(n, " reçus.")
				b = b[:n]
				file = append(file, b...)
				fmt.Println(b)
				lenght = lenght - int64(n)

				/*
					fmt.Println(lenght, " octets restants.")
					if lenght > util.MAXSIZE {
						b := make([]byte, util.MAXSIZE)
						_, err = s.Read(nfd, b) // Récupération d'un paquet de MAXSIZE octets
						check(err)
						fmt.Println(len(b))
						fmt.Println(b)
						file = append(file, b...)
						lenght = lenght - util.MAXSIZE
					} else {
						b := make([]byte, lenght)
						_, err = s.Read(nfd, b) // Récupération des derniers octets
						check(err)
						file = append(file, b...)
						break
					}*/
			}
			// Création du json
			//file = file[:len(file)-1]
			err = ioutil.WriteFile("./config.json", file, 0644)
			check(err)

			// Création de la structure
			clientListRep, err := util.BytesToFol(file)
			check(err)

			fmt.Println("Arborescence client :")
			fmt.Println(clientListRep.ToString())

			//err := ioutil.WriteFile("./test.txt", b, 0644)

		}(nfd, sa)
	}
}

func bigInt(size []byte, i int) int64 {
	if i < 7 {
		return int64(int32(size[i])<<uint8(56-i*8)) + bigInt(size, i+1)
	}
	return int64(size[i])
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
