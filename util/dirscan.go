package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Structure d'un fichier
type fic struct {
	nom     string
	lon     int64
	tim     time.Time
	md5hash []byte
}

// Structure d'un dossier
type fol struct {
	subFol []fol
	files  []fic
	nom    string
	tim    time.Time
}

// Calcul le MD5 d'un fichier
func calculerMd5(f *os.File) ([]byte, error) {
	hash := md5.New()
	_, err := io.Copy(hash, f)
	return hash.Sum(nil), err
}

// Convertie la structure fic sous forme de string
func (myFile fic) ToString() string {
	return "\n\t" + myFile.nom + "\t" + strconv.FormatInt(myFile.lon, 10) + "\t\t" + myFile.tim.Format("02/01/2006 15:04:05") + "\t\t" + fmt.Sprintf("%x", myFile.md5hash)
}

// Convertie la structure fpm sous forme de string
func (myFolder fol) ToString() string {
	var toPrint string = "\n" + myFolder.nom + "\t\t" + myFolder.tim.Format("02/01/2006 15:04:05")
	for _, fi := range myFolder.files {
		toPrint += fi.ToString()
	}
	for _, fo := range myFolder.subFol {
		toPrint += fo.ToString()
	}

	return toPrint
}

func openStat(file string) (fi *os.File, stat os.FileInfo, err error) {
	fi, err = os.Open(file)
	if err != nil {
		return
	}
	stat, err = fi.Stat()
	return
}

// Renvoie l'arborescence du dossier envoyé en paramètre en type fol
func ScanDir(folder string) (fol, error) {
	var listeRep fol

	// Récupération des infos sur le dossier actuel
	listeRep.nom = folder
	fo, stat, err := openStat(folder)
	defer fo.Close()
	if err != nil {
		return listeRep, err
	}
	listeRep.tim = stat.ModTime()

	// Parcours des fichiers et dossiers du dossier actuel
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return listeRep, err
	}
	for _, f := range files {
		curNam := f.Name()

		if f.IsDir() {

			// On ajoute un dossier dans le slice des dossiers
			fol, err := ScanDir(folder + curNam + "/")
			if err != nil {
				return listeRep, err
			}
			listeRep.subFol = append(listeRep.subFol, fol)
		} else {

			// On ajoute un fichier dans le slice des fichiers
			curSiz := f.Size()
			fi, _, err := openStat(folder + f.Name())
			defer fi.Close()
			if err != nil {
				return listeRep, err
			}
			hsh, err := calculerMd5(fi)
			if err != nil {
				return listeRep, err
			}
			listeRep.files = append(listeRep.files, fic{nom: curNam, lon: curSiz, tim: f.ModTime(), md5hash: hsh})
		}
	}
	return listeRep, err
}
