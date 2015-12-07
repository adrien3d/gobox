package util

import (
	"io/ioutil"
)

const (
	MAXSIZE = 500
)

// Cette fonction découpe un fichier en tableau de buffer
// de taille MAXSIZE.
func SplitFile(path string) (packets [][]byte, err error) {
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
