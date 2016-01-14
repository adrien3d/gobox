package util

import (
	"bytes"
	"crypto/md5"
	//"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
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

//  Entrée :
// - fol1 : dossier à comparer.
// - fol2 : dossier à mettre à jour.
//  Renvoie 2 structures Fol :
// - diff : contients tous les fichiers/dossiers qu'il faut importer sur dossier 2.
// - toDelete : contient tous les fichiers/dossiers qu'il faut supprimer sur dossier 1.
func CompareDir(fol1 Fol, fol2 Fol, lastUp time.Time) (Fol, Fol) {

	var diff, toDelete Fol
	diff.Nom = fol2.Nom
	if fol1.Tim.After(fol2.Tim) {
		diff.Tim = fol1.Tim
	} else {
		diff.Tim = fol2.Tim
	}

	for _, f1 := range fol1.Files {
		exist := false
		sameFile := 0
		for _, f2 := range fol2.Files {
			// si même nom
			if f1.Nom == f2.Nom {
				exist = true
				// et pas même fichier
				if !bytes.Equal(f1.Md5hash, f2.Md5hash) {
					if f1.Tim.After(f2.Tim) {
						diff.Files = append(diff.Files, f1) // on met à jour le fichier
					}
					break
				}
			} else {
				if bytes.Equal(f1.Md5hash, f2.Md5hash) {
					sameFile = sameFile + 1 // TODO : gérer les renommages
				}
			}
		}
		// n'existe pas
		if !exist {
			// et est plus récent
			if f1.Tim.After(lastUp) {
				diff.Files = append(diff.Files, f1) // on met à jour le fichier
			} else {
				// n'est pas plus récent et n'a pas le même contenu sous un autre nom
				if sameFile == 0 {
					toDelete.Files = append(toDelete.Files, f1) // fichier à supprimer sur f1
				}
				// TODO : gérer les renommages
			}
		}
	}

	for _, f1 := range fol1.SubFol {
		exist := false
		for _, f2 := range fol2.SubFol {
			// si même nom
			if lastFolder(f1.Nom) == lastFolder(f2.Nom) {
				tmp, tmp2 := CompareDir(f1, f2, lastUp)
				if len(tmp.Files) != 0 || len(tmp.SubFol) != 0 {
					diff.SubFol = append(diff.SubFol, tmp)
				}
				if len(tmp2.Files) != 0 || len(tmp2.SubFol) != 0 {
					toDelete.SubFol = append(toDelete.SubFol, tmp2)
				}
				exist = true
			}
		}
		// n'existe pas
		if !exist {
			if f1.Tim.After(lastUp) {
				diff.SubFol = append(diff.SubFol, f1) // on met à jour le dossier
			} else {
				toDelete.SubFol = append(toDelete.SubFol, f1) // dossier à supprimer sur f1
			}
		}
	}
	return diff, toDelete
}

func lastFolder(path string) string {
	folders := strings.Split(path, "/")
	return folders[len(folders)-2]
}
