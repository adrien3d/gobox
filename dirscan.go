package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "strconv"
)

type fic struct {
  nom string 
  lon int64
  tim time.Time
}
 
type fol struct {
  subFol []fol
  files  []fic
  nom string
  lon int64
  tim time.Time
}

func (myFile fic) fileToString() string {
  var toPrint string = "\n\t"+myFile.nom +"\t"+ strconv.FormatInt(myFile.lon, 10) +"\t\t"+ myFile.tim.Format("02/01/2006 15:04:05")
  return toPrint;
}

func (myFolder fol) folderToString() string {
  var toPrint string = "\n"+myFolder.nom +"\t\t"+ strconv.FormatInt(myFolder.lon, 10) +"\t\t"+ myFolder.tim.Format("02/01/2006 15:04:05")
 for _,fi := range myFolder.files {
      toPrint += fi.fileToString()
    }
  for _,fo := range myFolder.subFol {
    toPrint += fo.folderToString()
  }

  return toPrint;
}

func scanDir(folder string) (fol, error) {
	var listeRep fol
	listeRep.nom = folder

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		return listeRep, err
	}
	for _, f := range files {
		var curNam string = f.Name()

		fi, err := os.Open(folder + f.Name())
		defer fi.Close()
		if err != nil {
			return listeRep, err
		}
		stat, err := fi.Stat()
		if err != nil {
			return listeRep, err
		}
		var curSiz int64 = stat.Size()
    lastModifTime := stat.ModTime()

    listeRep.lon = curSiz
    listeRep.tim = lastModifTime

		if stat.IsDir() {
			var newPath string = folder + curNam + "/"

			fol, err := scanDir(newPath)
			if err != nil {
				return listeRep, err
			}
			listeRep.subFol = append(listeRep.subFol, fol)

		} else {
			fil := fic{nom: curNam, lon: curSiz, tim: lastModifTime}
			listeRep.files = append(listeRep.files, fil)
		}
	}
	return listeRep, err
}

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur fatale : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
  var dir string
  if (len(os.Args)>1) {
    dir = os.Args[1]
  } else {
    dir = "./files/"
  }
  listRep, err := scanDir(dir)
  check(err)
	fmt.Println(listRep.folderToString())
}
