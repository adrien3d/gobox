package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
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
	//var n int
	var conn util.Conn
	err := conn.Dial(PORT, ADDR)
	check(err)
	defer conn.Close()

	dat, err := util.SplitFile("./test.txt")
	check(err)
	err = conn.Write(dat[0])
	check(err)

	fmt.Println("FIN")
}

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
