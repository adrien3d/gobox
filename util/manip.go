package util

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	MAXPACKETSIZE = 1
)

// Lit le socket jusqu'à avoir un buffer de taille lenght
// Permet juste d'avoir une transaction de donnée plus stable !
func (c *Conn) Readbuffer(lenght int64) ([]byte, error) {
	// Création du buffer de réception
	maxsizepacket := int64(MAXPACKETSIZE)
	if lenght < MAXPACKETSIZE {
		maxsizepacket = lenght
	}
	file := make([]byte, 0, lenght+1)

	// Réception et assemblage
	for {
		//fmt.Println(lenght, " octets restants.")
		if lenght == 0 {
			return file, nil
		}
		b := make([]byte, maxsizepacket)
		n, err := c.Read(b)
		if err != nil {
			return file, err
		}
		//fmt.Println(n, " reçus.")
		b = b[:n]
		file = append(file, b...)
		lenght = lenght - int64(n)
	}
}

// Récupère un fichier :
// - D'abord récupère sa taille.
// - Puis récupère le fichier.
func (c *Conn) DownloadFile() ([]byte, error) {
	size, err := c.Readbuffer(8) // Réception de la taille du fichier en 8 octets
	if err != nil {
		return size, err
	}
	n := BigInt(size, 0)      // Conversion des 8 octets en entier
	err = c.Write([]byte{42}) // acknowledgment
	if err != nil {
		return size, err
	}
	tmp, err := c.Readbuffer(n) // Réception des n octets du fichier
	if err != nil {
		return tmp, err
	}
	err = c.Write([]byte{42}) // acknowledgment
	return tmp, err
}

// Envoi un fichier :
// - D'abord envoi sa taille.
// - Puis envoi le fichier.
func (c *Conn) UploadFile(path string) error {
	dat, err := SplitFile(path)
	if err != nil {
		return err
	}

	for _, packet := range dat {
		err = c.Write(packet)
		if err != nil {
			return err
		}
		_, err := c.Readbuffer(1) // acknowledgment
		if err != nil {
			return err
		}
	}
	return err
}

// Cette fonction découpe un fichier en tableau de packet de taille MAXSIZE.
// Le dernier packet est de taille < MAXSIZE.
//
// NB : le premier packet contient la taille du fichier (8 octets pour int64).
func SplitFile(path string) (packets [][]byte, err error) {
	absPath, _ := filepath.Abs(path)
	dat, err := ioutil.ReadFile(absPath)
	if err != nil {
		return
	}
	var lenght int = len(dat)
	packets = append(packets, Int64toByte(lenght))
	//packets = append(packets, strings.Fields(path))
	packets = append(packets, dat)
	return
}

func WriteFile(path string, buffer []byte) error {

	err := os.MkdirAll(folderOfFile(path), 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, buffer, 0644)
	return err
}

// Renvoie le tableau de byte de l'int64.
func Int64toByte(i int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

// Renvoie l'int16 du tableau de byte.
func BytetoInt(b []byte) int64 {
	i, _ := binary.Varint(b)
	return i
}

func BigInt(size []byte, i int) int64 {
	if i < 7 {
		return int64(int32(size[i])<<uint8(56-i*8)) + BigInt(size, i+1)
	}
	return int64(size[i])
}
