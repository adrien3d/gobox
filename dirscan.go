package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type fic struct {
	nom string
	lon int64
}

type fol struct {
	subFol []fol
	files  []fic
	nom    string
}

func scanDir(folder string) (fol, error) {
	var listeRep fol
	listeRep.nom = folder
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return listeRep, err
	}
	for _, f := range files {
		var curNam string = f.Name()

		fi, err := os.Open(folder + f.Name())
		defer fi.Close()
		if err != nil {
			return listeRep, err
		}
		stat, err := fi.Stat()
		if err != nil {

			return listeRep, err
		}
		var curSiz int64 = stat.Size()

		if stat.IsDir() {
			var newPath string = folder + curNam + "/"

			fol, err := scanDir(newPath)
			if err != nil {
				return listeRep, err
			}
			listeRep.subFol = append(listeRep.subFol, fol)

		} else {
			fil := fic{nom: curNam, lon: curSiz}
			listeRep.files = append(listeRep.files, fil)
		}
	}
	return listeRep, err
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	listRep, err := scanDir("./files/")
	check(err)
	fmt.Println(listRep)
}
