package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur fatale : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir = "./files/"
	}
	listRep, err := util.ScanDir(dir)
	check(err)
	fmt.Println(listRep.ToString())
}
