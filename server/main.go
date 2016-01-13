package main

import (
	"fmt"
	//"github.com/adrien3d/gobox/util"
	"io/ioutil"
	s "syscall"
)

const (
	PORT = 1002
)

func main() {

	// Mutexe de synchronisation
	/*
		var envoi = &sync.Mutex{}
		envoi.Lock()
		envoi.UnLock()*/

	// création du socket d'écoute

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

	// Lancement de l'écoute du serveur
	fmt.Println("Serveur lancé !")
	for {
		nfd, sa, err := s.Accept(fd)
		if err != nil {
			check(err)
		}

		go app(nfd, sa)
	}
}

func app(nfd int, sa s.Sockaddr) {
	fmt.Println("Nouveau socket :\n\t", sa)
	defer s.Close(nfd)
	lenght := bigInt(readbuffer(8, nfd), 0) // conversion des 8 octets en entier
	file := readbuffer(lenght, nfd)

	// Création du json
	err := ioutil.WriteFile("./config.json", file, 0644)
	check(err)

	// Création de la structure
	//clientListRep, err := util.BytesToFol(file)
	//check(err)

	//fmt.Println("Arborescence client :")
	//fmt.Println(clientListRep.ToString())

	//err := ioutil.WriteFile("./test.txt", b, 0644)

}
