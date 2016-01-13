package util

import (
	"encoding/binary"
	"io/ioutil"
	//"string"
)

const (
	// NB on est limité à 1350 octets via read/write
	MAXSIZE = 1000 // deprecated
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
	var lenght int = len(dat)
	packets = append(packets, int64toByte(lenght))
	//packets = append(packets, strings.Fields(path))
	packets = append(packets, dat)
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
