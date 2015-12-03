package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

type fic struct {
  nom string 
  lon int64
}
 
type fol struct {
  subFol []fol
  files  []fic
}

func scanDir(folder string) (fol) {
  var listeRep fol
      files, _ := ioutil.ReadDir(folder)
      for _, f := range files {
        var curNam string=f.Name()

        fi, _ := os.Open(f.Name())
        stat, _ := fi.Stat()
        var curSiz int64=stat.Size()

        if stat.IsDir(){
          s := []string{folder, curNam}
          var newPath string = strings.Join(s, "/")

          //fil := fic{nom: curNam, lon: curSiz}
          //fol := rep{subRep: scanDir(newPath), files: fil}
          //listeRep.subRep=append(listeRep.subRep, fol)

          fil := fic{nom: newPath, lon: curSiz}
          listeRep.files=append(listeRep.files, fil)

        } else {
          fil := fic{nom: curNam, lon: curSiz}
          listeRep.files=append(listeRep.files, fil)
        }
      }
  return listeRep
}

func checkErr(err error){
  if err != nil{
    fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
    os.Exit(1)
  }
}

func main() {
  listRep := scanDir("./")
  fmt.Println(listRep)
}