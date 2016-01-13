package main

import (
	"fmt"
	"os"
	s "syscall"
)

func readbuffer(lenght int64, nfd int) []byte {
	// Création du buffer de réception
	maxsizepacket := int64(1000)
	if lenght < 1000 {
		maxsizepacket = lenght
	}
	file := make([]byte, 0, lenght+1)
	//check(err)

	// Réception et assemblage
	for {
		fmt.Println(lenght, " octets restants.")
		if lenght == 0 {
			return file
		}

		b := make([]byte, maxsizepacket)
		n, err := s.Read(nfd, b)
		check(err)
		fmt.Println(n, " reçus.")
		b = b[:n]
		file = append(file, b...)
		lenght = lenght - int64(n)
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
