package util

import (
	"encoding/binary"
	"io/ioutil"
)

const (
	// NB on est limité à 1350 octets via read/write
	MAXSIZE = 1000
)

// Cette fonction découpe un fichier en tableau de packet de taille MAXSIZE.
// Le dernier packet est de taille < MAXSIZE.
//
// NB : le premier packet contient la taille du fichier (8 octets pour int64).
func SplitFile(path string) (packets [][]byte, err error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	packets = append(packets, int64toByte(len(dat)))
	for stop := false; stop == false; {
		var lenght int = len(dat)
		if lenght-MAXSIZE > 0 {
			packets = append(packets, dat[:MAXSIZE-1])
			dat = dat[MAXSIZE:lenght]
		} else {
			packets = append(packets, dat)
			stop = true
		}
	}
	return
}

// Renvoie le tableau de byte de l'int64.
func int64toByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

// Renvoie l'int16 du tableau de byte.
func BytetoInt(b []byte) int64 {
	i, _ := binary.Varint(b)
	return i
}
