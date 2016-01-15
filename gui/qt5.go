package main

import (
    "github.com/salviati/go-qt5/qt5"
)

func main() {
    qt5.Main(func() {
        w := qt5.NewWidget()
        w.SetWindowTitle(qt5.Version())
        w.SetSizev(300, 200)
        defer w.Close()
        w.Show()
        qt5.Run()
    })
}