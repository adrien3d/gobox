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

//  Entrée :
// - fol1 : dossier à comparer.
// - fol2 : dossier à mettre à jour.
//  Renvoie 2 structures Fol :
// - upName : contient tous les fichiers/dossiers qu'il faut renommer sur dossier 2.
// - diff : contients tous les fichiers/dossiers qu'il faut importer sur dossier 2.
func CompareDir(fol1 Fol, fol2 Fol) Fol {
	var diff Fol
	var upName Fol
	for {
		// Si le dossier 1 est vide il n'y a plus de fichier à mettre dans diff
		if len(fol1.Files) == 0 {
			break
		}

		// Si le dossier 2 est vide tous le dossier client va dans diff
		if len(fol2.Files) == 0 {
			for {
				f1, fol1.Files = fol1.Files[len(fol1.Files)-1], fol1.Files[:len(fol1.Files)-1]
				diff.Files = append(diff.Files, f1)
				if len(fol1.Files) == 0 {
					break
				}
			}
			break
		}

		// on pop 1 élément de chaque dossier
		f1, fol1.Files = fol1.Files[len(fol1.Files)-1], fol1.Files[:len(fol1.Files)-1]
		f2, fol2.Files = fol2.Files[len(fol2.Files)-1], fol2.Files[:len(fol2.Files)-1]

		// Si il y a deux contenus différents
		if f1.Md5hash != f2.Md5hash {
			// mais de même nom
			if f1.Nom == f2.Nom {
				// et que le fichier 1 est plus récent
				if f1.Tim.After(f2.Tim) {
					diff.Files = append(diff.Files, f1) // on met à jour le fichier
				}
				// Sinon on a déjà la version la plus récente
			} else {
				exist := false
				for _, f2 := range fol2.Files {
					// Si même contenu
					if f1.Md5hash == f2.Md5hash {
						// mais différent nom
						if f1.Nom != f2.Nom {
							// et que le fichier 1 est plus récent
							if f1.Tim.After(f2.Tim) {
								upName.Files = append(upName.Files, f1) // on met à jour le nom
							}
						}
						exist = true
						break
					}
					// Si même nom
					if f1.Nom == f2.Nom {
						// et que le fichier 1 est plus récent
						if f1.Tim.After(f2.Tim) {
							diff.Files = append(diff.Files, f1) // on met à jour le fichier
						}
						// Sinon on a déjà la version la plus récente
					}
				}
				if !exist {
					// contenu nouveau et nom nouveau
					diff.Files = append(diff.Files, f1)
				}
			}
		} else {
			// Si même contenu et différent nom
			if f1.Nom != f2.Nom {
				// et que le fichier 1 est plus récent
				if f1.Tim.After(f2.Tim) {
					upName.Files = append(upName.Files, f1) // on met à jour le nom
				}
			}
		}

	}

}

func CompareDirRec(fol1 Fol, fol2 Fol, fol3 Fol) {

}
