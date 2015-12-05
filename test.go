package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"io/ioutil"
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

	var listRep util.Fol
	err := util.ScanDir(dir, &listRep)
	check(err)
	//fmt.Println(listRep.ToString())
	//fmt.Println(listRep)
	b, err := listRep.ToBytes()
	//b, err := json.Marshal(listRep)
	check(err)
	err = ioutil.WriteFile("./test.json", b, 0644)
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
}
