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

/*type Fic struct {
	Nom     string    `json:"fileName"`
	Lon     int64     `json:"size"`
	Tim     time.Time `json:"lastFileUpdate"`
	Md5hash []byte    `json:"md5"`
}

type Fol struct {
	SubFol []Fol     `json:"listFolders"`
	Files  []Fic     `json:"listFiles"`
	Nom    string    `json:"folderName"`
	Tim    time.Time `json:"lastFolderUpdate"`
}*/
// Renvoie l'arborescence du dossier envoyé en paramètre en type Fol
func ScanDir(folder string, listeRep *Fol) error {
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
			fo, stat, err := openStat(folder + curNam)
			defer fo.Close()
			if err != nil {
				return err
			}
			if stat.ModTime().After(listeRep.Tim) {
				listeRep.Tim = stat.ModTime()
			}
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

/*Comparer la structure des dossiers
  Éliminer les doublons, en se basant sur le md5
  Si md5 différent, garder le ficher avec la date la plus récente
func CompareDir(fol1 Fol, fol2 Fol) Fol {
	var difFol Fol
	for _, f1 := range fol1 {
		for _, f2 := range fol2 {
			if f1.(Fic) && f2.(Fic) {
				if f1.Nom == f2.Nom {
					if f1.Md5hash != f2.Md5hash {
						if f1.Tim.After(f2.Tim) {
							difFol.Files = append(difFol.Files, f1)
						} else  {
							difFol.Files = append(difFol.Files, f2)
						}
					}
				}
			}
		}
	}
	return difFol
}
*/
