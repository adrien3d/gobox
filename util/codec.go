package util

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// Structure d'un Fichier
type Fic struct {
	Nom     string    `json:"fileName"`
	Lon     int64     `json:"size"`
	Tim     time.Time `json:"lastFileUpdate"`
	Md5hash []byte    `json:"md5"`
}

// Structure d'un dossier
type Fol struct {
	SubFol []Fol     `json:"listFolders"`
	Files  []Fic     `json:"listFiles"`
	Nom    string    `json:"folderName"`
	Tim    time.Time `json:"lastFolderUpdate"`
}

func (myFol Fol) ToJson() ([]byte, error) {
	return json.MarshalIndent(myFol, "", " ")
}

func (myFol Fol) ToBytes() ([]byte, error) {
	return json.Marshal(myFol)
}

func BytesToFol(b []byte) (myFol Fol, err error) {
	err = json.Unmarshal(b, &myFol)
	return
}

// Convertit la structure Fic sous forme de string
func (myFile Fic) ToString() string {
	return "\n\t" + myFile.Nom + "\t" + strconv.FormatInt(myFile.Lon, 10) + "\t\t" + myFile.Tim.Format("02/01/2006 15:04:05") + "\t\t" + fmt.Sprintf("%x", myFile.Md5hash)
}

// Convertit la structure Fol sous forme de string
func (myFolder Fol) ToString() string {
	var toPrint string = "\n" + myFolder.Nom + "\t\t" + myFolder.Tim.Format("02/01/2006 15:04:05")
	for _, fi := range myFolder.Files {
		toPrint += fi.ToString()
	}
	for _, fo := range myFolder.SubFol {
		toPrint += fo.ToString()
	}
	return toPrint
}

func (myFol Fol) Parcours() []Fic {

	for i := range myFol.Files {
		myFol.Files[i].Nom = myFol.Nom + myFol.Files[i].Nom
	}
	lesFics := myFol.Files
	for _, subfol := range myFol.SubFol {
		lesFics = append(lesFics, subfol.Parcours()...)
	}

	return lesFics
}
