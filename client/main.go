package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
)

const (
	PORT   = 1002
	FOLDER = "./gobox/"
)

var (
	//ADDR = [4]byte{10, 8, 0, 1}
	//ADDR = [4]byte{127, 0, 0, 1}
	ADDR = [4]byte{5, 39, 89, 231}
)

func main() {

	// Scan du répertoire à synchroniser
	var listRep util.Fol
	err := util.ScanDir(FOLDER, &listRep)
	check(err)
	b, err := listRep.ToBytes()
	check(err)

	// Etablissement de la connexion au serveur
	var conn util.Conn
	err = conn.Dial(PORT, ADDR)
	check(err)
	defer conn.Close()

	// Attente d'une réponse serveur
	ack, err := conn.Readbuffer(1)
	fmt.Println(ack[0])

	// Envoi de l'arborescence sous forme de Json
	fmt.Printf("Envoi de l'arborescence")
	err = conn.Write(util.Int64toByte(len(b)))
	check(err)
	err = conn.Write(b)
	check(err)

	// Réception du calcul des différences
	fmt.Printf("Réception du calcul des différences...")

	// diff1
	tmp, err := conn.DownloadFile() // téléchargement d'un fichier
	check(err)
	diff1, err := util.BytesToFol(tmp) // conversion en structure Fol
	check(err)
	toSend := diff1.Parcours() // tableau des fichiers à envoyer

	// diff2
	tmp2, err := conn.DownloadFile()
	check(err)
	diff2, err := util.BytesToFol(tmp2)
	check(err)
	toGet := diff2.Parcours()

	// del 1
	tmp3, err := conn.DownloadFile()
	check(err)
	del, err := util.BytesToFol(tmp3)
	check(err)
	toDel := del.Parcours()

	// Suppression des fichiers locaux
	for _, file := range toDel {
		check(os.Remove(file.Nom))
	}

	// Envoi des fichiers client vers serveur
	for _, file := range toSend {
		err = conn.UploadFile(file.Nom)
		check(err)
	}

	// Réception des fichiers serveur vers client
	for _, file := range toGet {
		newfile, err := conn.DownloadFile()
		check(err)
		check(ioutil.WriteFile(file.Nom, newfile, 0644))
	}

	fmt.Println("Synchronisation effectuée avec succès.")
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
