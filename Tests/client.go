package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "10.8.0.1:3000")
	check(err)

	dat, err := ioutil.ReadFile("./test.txt")
	check(err)
	//message := bufio.NewWriter(conn).Write(dat)
	conn.Write(dat)
	//fmt.Println(message)

}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
