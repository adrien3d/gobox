package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "io"
    "crypto/md5"
    "time"
    "strconv"
)

type fic struct {
  nom string 
  lon int64
  tim time.Time
  md5hash []byte
}
 
type fol struct {
  subFol []fol
  files  []fic
  nom string
  tim time.Time
}

func calculerMd5(filePath string) ([]byte, error) {
  var result []byte
  file, err := os.Open(filePath)
  if err != nil {
    return result, err
  }
  defer file.Close()
 
  hash := md5.New()
  if _, err := io.Copy(hash, file); err != nil {
    return result, err
  }
 
  return hash.Sum(result), nil
}

func (myFile fic) fileToString() string {
  var toPrint string = "\n\t"+myFile.nom +"\t"+ strconv.FormatInt(myFile.lon, 10) +"\t\t"+ myFile.tim.Format("02/01/2006 15:04:05")+"\t\t"+ string(myFile.md5hash[:])
  return toPrint;
}

func (myFolder fol) folderToString() string {
  var toPrint string = "\n"+myFolder.nom +"\t\t"+ myFolder.tim.Format("02/01/2006 15:04:05")
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

    listeRep.tim = lastModifTime

    var newPath string = folder + curNam + "/"

		if stat.IsDir() {
			fol, err := scanDir(newPath)
			if err != nil {
				return listeRep, err
			}
			listeRep.subFol = append(listeRep.subFol, fol)
		} else {
      hsh, _ := calculerMd5(curNam); 
      fmt.Printf("\nmd5 checksum is: %x", hsh)
			fil := fic{nom: curNam, lon: curSiz, tim: lastModifTime, md5hash:hsh}
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

  /*if b, err := calculerMd5("dirscan.go"); err != nil {
    fmt.Printf("Err: %v", err)
  } else {
    fmt.Printf("main.go md5 checksum is: %x", b)
  }*/
}
