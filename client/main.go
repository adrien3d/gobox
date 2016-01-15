package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"os"
)

const (
	PORT   = 1002
	FOLDER = "./gobox/"
)

var (
	//ADDR = [4]byte{10, 8, 0, 1}
	ADDR = [4]byte{127, 0, 0, 1}
	//ADDR = [4]byte{5, 39, 89, 231}
)

func main() {
	fmt.Printf("Demarrage client")
	check(os.MkdirAll(FOLDER, 0777))

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
	for {
		ack, err := conn.Readbuffer(1)
		check(err)
		if len(ack) == 1 && ack[0] == 42 {
			break
		}
	}

	// Envoi de l'arborescence sous forme de Json
	fmt.Printf("\nEnvoi de l'arborescence\n\n")
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

	fmt.Println(toSend)

	// diff2
	tmp2, err := conn.DownloadFile()
	check(err)
	diff2, err := util.BytesToFol(tmp2)
	check(err)
	toGet := diff2.Parcours()

	fmt.Println(toGet)

	// del 1
	tmp3, err := conn.DownloadFile()
	check(err)
	del, err := util.BytesToFol(tmp3)
	check(err)
	toDel := del.Parcours()

	fmt.Println(toDel)

	// Suppression des fichiers locaux
	/*for _, file := range toDel {
		fmt.Println("Suppression de ", file.Nom)
		check(util.DeleteFile(file.Nom))
	}*/

	// Envoi des fichiers client vers serveur
	for _, file := range toSend {
		fmt.Println("Envoi de ", file.Nom)
		err = conn.UploadFile(file.Nom)
		check(err)
	}

	// Réception des fichiers serveur vers client
	for _, file := range toGet {
		fmt.Println("Reception de ", file.Nom)
		newfile, err := conn.DownloadFile()
		check(err)
		check(util.WriteFile(file.Nom, newfile))
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
