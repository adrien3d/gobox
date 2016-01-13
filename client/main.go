package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
	"sync"
)

const (
	PORT = 1002
)

var (
	//ADDR = [4]byte{10, 8, 0, 1}
	ADDR = [4]byte{5, 39, 89, 231}
)

func main() {

	// Etablissement de la connexion au serveur
	var conn util.Conn
	err := conn.Dial(PORT, ADDR)
	check(err)
	defer conn.Close()

	// Scan du répertoire à synchroniser
	var listRep util.Fol
	err = util.ScanDir("../../", &listRep)
	check(err)
	b, err := listRep.ToBytes()
	check(err)
	err = ioutil.WriteFile("./test.json", b, 0644)
	check(err)
	dat, err := util.SplitFile("./test.json")

	// Envoi des packets d'un fichier

	for i, packet := range dat {
		fmt.Printf("Envoi du packet N°%d.\n", i)
		fmt.Printf("\t%d octets envoyés ", len(packet))
		err = conn.Write(packet)
		check(err)
		fmt.Printf("avec succès.")
	}
	fmt.Println("FIN")
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
