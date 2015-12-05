package util

import (
	"crypto/md5"
	"io"
	"io/ioutil"
	"os"
)

// Calcul le MD5 d'un Fichier
func calculerMd5(f *os.File) ([]byte, error) {
	hash := md5.New()
	_, err := io.Copy(hash, f)
	return hash.Sum(nil), err
}

func openStat(file string) (fi *os.File, stat os.FileInfo, err error) {
	fi, err = os.Open(file)
	if err != nil {
		return
	}
	stat, err = fi.Stat()
	return
}

// Renvoie l'arborescence du dossier envoyé en paramètre en type Fol
func ScanDir(folder string, listeRep *Fol) error {
	//var listeRep Fol

	// Récupération des infos sur le dossier actuel
	listeRep.Nom = folder
	fo, stat, err := openStat(folder)
	defer fo.Close()
	if err != nil {
		return err
	}
	listeRep.Tim = stat.ModTime()

	// Parcours des Fichiers et dossiers du dossier actuel
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}
	for _, f := range files {
		curNam := f.Name()

		if f.IsDir() {

			// On ajoute un dossier dans le slice des dossiers
			var fol Fol
			err := ScanDir(folder+curNam+"/", &fol)
			if err != nil {
				return err
			}
			listeRep.SubFol = append(listeRep.SubFol, fol)
		} else {

			// On ajoute un Fichier dans le slice des Fichiers
			curSiz := f.Size()
			fi, _, err := openStat(folder + f.Name())
			defer fi.Close()
			if err != nil {
				return err
			}
			hsh, err := calculerMd5(fi)
			if err != nil {
				return err
			}
			listeRep.Files = append(listeRep.Files, Fic{Nom: curNam, Lon: curSiz, Tim: f.ModTime(), Md5hash: hsh})
		}
	}
	return err
}
