package main

import (
	"encoding/json"
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
	"sync"
	s "syscall"
	"time"
)

const (
	PORT         = 1002
	FOLDER       = "./gobox/"
	LASTUPFOLDER = "./lastup.json"
)

var (
	lastUpdate time.Time
	envoi      *sync.Mutex
	mainFolder util.Fol
)

func main() {
	check(os.MkdirAll(FOLDER, 0777))

	// Initialisation des variables serveur
	lastUpdate = time.Date(1994, time.April, 2, 2, 0, 0, 0, time.UTC)
	dat, err := ioutil.ReadFile(LASTUPFOLDER)
	if err == nil {
		err = json.Unmarshal(dat, &lastUpdate)
		check(err)
	}

	envoi = &sync.Mutex{}
	err = util.ScanDir(FOLDER, &mainFolder)
	check(err)

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
		envoi.Lock()
		go app(nfd, sa)
		envoi.Unlock()
	}
}

func app(nfd int, sa s.Sockaddr) {
	defer s.Close(nfd)
	//inet4 := sa.sockaddr().
	//s.SockaddrInet4{Port: sa.Port, Addr: sa.Addr}
	fmt.Printf("%v", sa)
	conn := util.SetConn(nfd)

	// Envoie de l'acknowledge pour lancer la synchro
	for {
		if conn.Write([]byte{42}) == nil {
			break
		}
	}

	fmt.Println("Nouveau socket :\n\t", sa)
	size, err := conn.Readbuffer(8)
	lenght := util.BigInt(size, 0) // conversion des 8 octets en entier
	file, err := conn.Readbuffer(lenght)

	// Création du json
	//err := ioutil.WriteFile("./config.json", file, 0644)
	//check(err)

	// Création de la structure
	clientListRep, err := util.BytesToFol(file)
	check(err)

	// Calcul des différences
	diff2, del1 := util.CompareDir(clientListRep, mainFolder, lastUpdate)
	diff1, del2 := util.CompareDir(mainFolder, clientListRep, lastUpdate)
	fmt.Println("\n*** Calcul des fichiers à mettre à jour client vers serveur ***")
	fmt.Println(diff2)
	fmt.Println("\n*** Fichiers à supprimer sur le serveur ***")
	fmt.Println(del2)
	fmt.Println("\n*** Calcul des fichiers à mettre à jour serveur vers client ***")
	fmt.Println(diff1)
	fmt.Println("\n*** Fichiers à supprimer sur le client ***")
	fmt.Println(del1)

	//toDel := del2.Parcours()

	// Envoi de la structure diff2
	buff1, err := diff2.ToBytes()
	check(err)
	err = conn.Write(util.Int64toByte(len(buff1)))
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	err = conn.Write(buff1)
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	toGet := diff2.Parcours()
	fmt.Println("Diff2 envoyé")

	// Envoi de la structure diff1
	buff2, err := diff1.ToBytes()
	check(err)
	err = conn.Write(util.Int64toByte(len(buff2)))
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	err = conn.Write(buff2)
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	toSend := diff1.Parcours()
	fmt.Println("Diff1 envoyé")

	// Envoi de la structure del1
	buff3, err := del1.ToBytes()
	check(err)
	err = conn.Write(util.Int64toByte(len(buff3)))
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	err = conn.Write(buff3)
	check(err)
	_, err = conn.Readbuffer(1) // acknowledgment
	check(err)
	fmt.Println("Del1 envoyé")

	// Suppression des fichiers del2
	/*for _, file := range toDel {
		fmt.Println("Suppression de ", file.Nom)
		check(util.DeleteFile(file.Nom))
	}*/

	// Réception des fichiers diff2
	for _, file := range toGet {
		fmt.Println("Reception de ", file.Nom)
		newfile, err := conn.DownloadFile()
		check(err)
		check(util.WriteFile(file.Nom, newfile))
	}

	// Envoie des fichiers diff1
	for _, file := range toSend {
		fmt.Println("Envoi de ", file.Nom)
		err = conn.UploadFile(file.Nom)
		check(err)
	}

	// mise à jour de l'heure de dernière synchronisation
	lastUpdate = time.Now()
	currentTime, err := json.Marshal(lastUpdate)
	check(err)
	check(util.WriteFile(LASTUPFOLDER, currentTime))

	fmt.Println("Synchronisation effectuée avec succès !")

}
