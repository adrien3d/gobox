package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	MAXSIZE = 500
)

func main() {
	//serverAddr, err := net.ResolveTCPAddr("tcp", "5.39.89.231:3000")
	//conn, err := net.DialTCP("tcp", nil, serverAddr)
	//conn, err := net.Dial("tcp", "10.8.0.1:3000")
	conn, err := net.Dial("tcp", "5.39.89.231:3000")
	check(err)

	dat, err := splitFile("./test.txt")
	fmt.Println(dat)
	check(err)

	conn.Write(dat[0])

}

func splitFile(path string) (packets [][]byte, err error) {

	dat, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		return
	}
	for stop := false; stop == false; {

		if len(dat)-MAXSIZE > 0 {
			packets = append(packets, dat[:MAXSIZE-1])
			dat = dat[MAXSIZE:len(dat)]
		} else {
			packets = append(packets, dat)
			stop = true
		}
	}

	return
}
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
