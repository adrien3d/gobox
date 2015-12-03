package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	s "syscall"
)

const (
	MAXSIZE = 500
	PORT    = 3000
)

var (
	ADDR = [4]byte{10, 8, 0, 1}
)

// Etablie une connexion via socket sur le serveur
// à l'aide des librairies POSIX.
func Dial() (s.Handle, s.SockaddrInet4, error) {
	var sa s.SockaddrInet4 = s.SockaddrInet4{Port: PORT, Addr: ADDR}
	var d s.WSAData
	var sd s.Handle
	if runtime.GOOS == "windows" {
		err := s.WSAStartup(uint32(0x202), &d)
		if err != nil {
			return sd, sa, err
		}
	}

	sd, err := s.Socket(s.AF_INET, s.SOCK_STREAM, 0)
	if err != nil {
		return sd, sa, err
	}
	//if err := s.Bind(sd, &sa); err != nil {
	//	return sd, err
	//}
	s.Connect(sd, &sa)
	return sd, sa, err
}

func main() {
	//var n int
	sd, sa, err := Dial()
	check(err)
	defer s.Close(sd)

	dat, err := splitFile("./test.txt")
	check(err)

	if runtime.GOOS == "windows" {
		data := dat[0]
		buf := &s.WSABuf{
			Len: uint32(len(data)),
			Buf: &data[0],
		}
		var sent *uint32
		overlapped := s.Overlapped{}
		croutine := byte(0)

		err = s.WSASendto(sd, buf, 1, sent, uint32(0), &sa, &overlapped, &croutine)
	} else {
		n, err = s.Write(sd, dat[0])
	}
	check(err)

	fmt.Println("FIN")
}

// Cette fonction découpe un fichier en tableau de buffer
// de taille MAXSIZE.
func splitFile(path string) (packets [][]byte, err error) {

	dat, err := ioutil.ReadFile(path)
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

// Fonction pour checker les erreurs
func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
