package main

import (
	"encoding/json"
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
	"os"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur fatale : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	/*var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "./files/"
	}*/

	lastUpdate := time.Date(2016, time.January, 13, 12, 0, 0, 0, time.UTC)
	dat, err := ioutil.ReadFile("./lastUp.json")
	if err == nil {
		err = json.Unmarshal(dat, &lastUpdate)
		check(err)
	}

	var listRepClient, listRepServer util.Fol
	dirClient := "./files/"
	dirServer := "./files2/"
	err = util.ScanDir(dirClient, &listRepClient)
	check(err)
	err = util.ScanDir(dirServer, &listRepServer)
	check(err)
	/*fmt.Println("\n*** Affichage du client ***")
	fmt.Println(listRepClient.ToString())
	fmt.Println("\n*** Affichage du serveur ***")
	fmt.Println(listRepServer.ToString())*/
	diff2, del1 := util.CompareDir(listRepClient, listRepServer, lastUpdate)
	diff1, del2 := util.CompareDir(listRepServer, listRepClient, lastUpdate)
	fmt.Println("\n*** Calcul des fichiers à mettre à jour client vers serveur ***")
	fmt.Println(diff2.ToString())
	fmt.Println("\n*** Fichiers à supprimer sur le serveur ***")
	fmt.Println(del2.ToString())
	fmt.Println("\n*** Calcul des fichiers à mettre à jour serveur vers client ***")
	fmt.Println(diff1.ToString())
	fmt.Println("\n*** Fichiers à supprimer sur le client ***")
	fmt.Println(del1.ToString())
	b, err := json.Marshal(time.Now())
	check(err)
	err = ioutil.WriteFile("./lastUp.json", b, 0644)
	check(err)
	//fmt.Println("\n*** Calcul des noms à mettre à jour ***")

	//fmt.Println(upName.ToString())
	/*
		b, err := listRep.ToBytes() //Conversion en JSON : marshallification
		check(err)
		err = ioutil.WriteFile("./test.json", b, 0644) //Sauvegarde dans le fichier
		check(err)
		fi, err := os.Open("./test.json")
		defer fi.Close()
		stat, _ := fi.Stat()
		size := stat.Size()
		c := make([]byte, size)
		fi.Read(c)
		newListRep, err := util.BytesToFol(c)
		check(err)
		fmt.Println(newListRep.ToString())
	*/
}
