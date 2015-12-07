package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
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
	err = util.ScanDir("../files/", &listRep)
	check(err)
	b, err := listRep.ToBytes()
	check(err)
	err = ioutil.WriteFile("./test.json", b, 0644)
	check(err)
	err = conn.Write(b)
	check(err)

	//dat, err := util.SplitFile("./test.txt")
	fmt.Println("FIN")
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
