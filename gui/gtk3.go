package main

import (
    "github.com/gotk3/gotk3/gtk"
    "log"
)

func main() {
    gtk.Init(nil)

    win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
    if err != nil {
        log.Fatal("Impossible de créer la fenêtre :", err)
    }
    win.SetTitle("GoBox a0.1")
    win.Connect("destroy", func() {
        gtk.MainQuit()
    })

    grid, err := gtk.GridNew()
    if err != nil {
        log.Fatal("Impossible de créer la grille :", err)
    }


    label1, err := gtk.LabelNew("Adresse IP / Port : ")
    if err != nil {
        log.Fatal("Impossible de créer le label IP :", err)
    }
    
    label3, err := gtk.LabelNew("Dossier à synchroniser : ")
    if err != nil {
        log.Fatal("Impossible de créer le label Dossier :", err)
    }

    entry1, err := gtk.EntryNew()
    if err != nil {
        log.Fatal("Impossible de créer le champ IP :", err)
    }    
    entry2, err := gtk.EntryNew()
    if err != nil {
        log.Fatal("Impossible de créer le champ Port :", err)
    }    
    entry3, err := gtk.EntryNew()
    if err != nil {
        log.Fatal("Impossible de créer le champ Dossier :", err)
    }

    btn, err := gtk.ButtonNewWithLabel("Lancer la synchronisation")
    if err != nil {
        log.Fatal("Impossible de créer le bouton synchronisation :", err)
    }

    /*btn2, err := gtk.FileChooserButtonNew("Choix")
    if err != nil {
        log.Fatal("Impossible de créer le bouton choix :", err)
    }*/

    grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
//Attach(child IWidget, left, top, width, height int)
    grid.Add(label1)
    grid.SetOrientation(gtk.ORIENTATION_HORIZONTAL)
    grid.Add(entry1)
    grid.Add(entry2)


    grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
    grid.Add(label3)
    grid.Add(entry3)

    grid.Attach(btn, 1, 2, 1, 2)

    btn.Connect("clicked", func() {
        /*dialog, _ := gtk.DialogNew()

        filechooser, _ := gtk.FileChooserWidgetNew(gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
        dialog.AddButton("Valider", gtk.RESPONSE_ACCEPT)
        dialog.AddButton("Annuler", gtk.RESPONSE_CANCEL)
        dialog.SetTitle("Choisir le dossier a synchroniser")
        box, _ := dialog.GetContentArea()
        box.Add(filechooser)
        box.ShowAll()
        log.Print("Clic lancer synchro")*/
        filechooserdialog, _ := gtk.FileChooserDialogNewWith1Button(
            "Choisissez un fichier ...",
            //btn.GetTopLevelAsWindow(),
            win,
            gtk.FILE_CHOOSER_ACTION_OPEN,
            "Valider",
            gtk.RESPONSE_ACCEPT)
        /*filter := gtk.NewFileFilter()
        filter.AddPattern("*.go")
        filechooserdialog.AddFilter(filter)*/
        filechooserdialog.Response(func() {
            println(filechooserdialog.GetFilename())
            filechooserdialog.Destroy()
        })
        filechooserdialog.Run()
    })
    /*
    nbChildAll, err := gtk.LabelNew("Tous mes fichiers sont ici")
    if err != nil {
        log.Fatal("Unable to create button:", err)
    }
    nbTabAll, err := gtk.LabelNew("Tout")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }

    nbChildMusic, err := gtk.LabelNew("Toute ma musique est ici")
    if err != nil {
        log.Fatal("Unable to create button:", err)
    }
    nbTabMusic, err := gtk.LabelNew("Musique")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }

    nbChildPhotos, err := gtk.LabelNew("Toutes mes photos sont ici")
    if err != nil {
        log.Fatal("Unable to create button:", err)
    }
    nbTabPhotos, err := gtk.LabelNew("Photos")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }

    nbChildVideos, err := gtk.LabelNew("Toutes mes vidéos sont ici")
    if err != nil {
        log.Fatal("Unable to create button:", err)
    }
    nbTabVideos, err := gtk.LabelNew("Vidéos")
    if err != nil {
        log.Fatal("Unable to create label:", err)
    }
    nb.AppendPage(nbChildAll, nbTabAll)
    nb.AppendPage(nbChildMusic, nbTabMusic)
    nb.AppendPage(nbChildPhotos, nbTabPhotos)
    nb.AppendPage(nbChildVideos, nbTabVideos)*/

    win.Add(grid)
    win.SetDefaultSize(200, 250)
    win.ShowAll()

    gtk.Main()
}
