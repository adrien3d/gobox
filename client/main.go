package main

import (
	"fmt"
	"github.com/adrien3d/gobox/util"
	"os"
	s "syscall"
)

func main() {
	//var n int
	sd, sa, err := util.Dial()
	check(err)
	defer s.Close(sd)

	dat, err := util.SplitFile("./Tests/test.txt")
	check(err)
	err = util.Write(sd, sa, dat[0])
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
