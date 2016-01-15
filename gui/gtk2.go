package main

import (
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
    /*"fmt"
    "log"
    "strings"
    "bytes"*/
    //"os"
    //"strconv"
    "os/exec"
)

func sync(folder, address, port string) {
    /*port2, _ :=strconv.Atoi(port)
    addr := [4]byte{5, 39, 89, 231}*/
    
    //folder = "../client/gobox/"

    /*mes4octets := strings.Split(taString, ".")
    addr[0] = byte(mes4octets[0])*/
    println("lancer client")

    //println(exec.Command("./client"/*, address, port*/))
    exec.Command("/bin/sh","./client")
    /*cmd := exec.Command("./client")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("in all caps: %q\n", out.String())*/
}

func main() {
    var menuitem *gtk.MenuItem
    gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetPosition(gtk.WIN_POS_CENTER)
    window.SetTitle("GoBox a0.1")
    window.SetIconName("gtk-dialog-info")
    window.Connect("destroy", func(ctx *glib.CallbackContext) {
        println("got destroy!", ctx.Data().(string))
        gtk.MainQuit()
    }, "foo")

    //--------------------------------------------------------
    // GtkVBox
    //--------------------------------------------------------
    vbox := gtk.NewVBox(false, 1)

    //--------------------------------------------------------
    // GtkMenuBar
    //--------------------------------------------------------
    menubar := gtk.NewMenuBar()
    vbox.PackStart(menubar, false, false, 0)

    //--------------------------------------------------------
    // GtkVPaned
    //--------------------------------------------------------
    vpaned := gtk.NewVPaned()
    vbox.Add(vpaned)

    //--------------------------------------------------------
    // GtkFrame
    //--------------------------------------------------------
    frame1 := gtk.NewFrame("Dossier et Paramètres")
    framebox1 := gtk.NewVBox(false, 1)
    frame1.Add(framebox1)

    frame2 := gtk.NewFrame("Fonctions")
    framebox2 := gtk.NewVBox(false, 1)
    frame2.Add(framebox2)

    vpaned.Pack1(frame1, false, false)
    vpaned.Pack2(frame2, false, false)

    //--------------------------------------------------------
    // GtkImage
    //--------------------------------------------------------
    /*dir, _ := path.Split(os.Args[0])
    //imagefile := path.Join(dir, "../../mattn/go-gtk/data/go-gtk-logo.png")
    imagefile := path.Join(dir, "./go-gtk-logo.png")
    println(dir)*/

    label := gtk.NewLabel("GoBox a0.1")
    label.ModifyFontEasy("DejaVu Serif 15")
    framebox1.PackStart(label, false, true, 0)

    //--------------------------------------------------------
    // GtkEntry
    //--------------------------------------------------------
    champIp := gtk.NewEntry()
    champIp.SetText("10.0.0.1")
    framebox1.Add(champIp)    

    champPort := gtk.NewEntry()
    champPort.SetText("80")
    framebox1.Add(champPort)

    folder := "./"

    /*image := gtk.NewImageFromFile(imagefile)
    framebox1.Add(image)*/
    buttons := gtk.NewHBox(false, 1)
    //--------------------------------------------------------
    // GtkButton
    //--------------------------------------------------------
    button := gtk.NewButtonWithLabel("Choisir le dossier")
    button.Clicked(func() {
            //--------------------------------------------------------
            // GtkFileChooserDialog
            //--------------------------------------------------------
            filechooserdialog := gtk.NewFileChooserDialog(
                "Sélectionnez le dossier ...",
                button.GetTopLevelAsWindow(),
                gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
                gtk.STOCK_OK,
                gtk.RESPONSE_ACCEPT)
            /*filter := gtk.NewFileFilter()
            filter.AddPattern("*.go")
            filechooserdialog.AddFilter(filter)*/
            filechooserdialog.Response(func() {
                println(filechooserdialog.GetFilename())
                folder = filechooserdialog.GetFilename() + "/"
                filechooserdialog.Destroy()
            })
            filechooserdialog.Run()
    })
    buttons.Add(button)

    
    //--------------------------------------------------------
    // GtkToggleButton
    //--------------------------------------------------------
    togglebutton := gtk.NewToggleButtonWithLabel("Lancer la synchronisation")
    togglebutton.Connect("toggled", func() {
        if togglebutton.GetActive() {
            togglebutton.SetLabel("Synchronisation ON")
            //Appel fonction synchro avec paramètres
            println(folder, champIp.GetText(), champPort.GetText())
            sync(folder, champIp.GetText(), champPort.GetText())
        } else {
            togglebutton.SetLabel("Synchronisation OFF")
        }
    })
    buttons.Add(togglebutton)

    framebox2.PackStart(buttons, false, false, 0)

    //--------------------------------------------------------
    // GtkVSeparator
    //--------------------------------------------------------
    vsep := gtk.NewVSeparator()
    framebox2.PackStart(vsep, false, false, 0)


    //--------------------------------------------------------
    // GtkMenuItem
    //--------------------------------------------------------
    cascademenu := gtk.NewMenuItemWithMnemonic("_Fichier")
    menubar.Append(cascademenu)
    submenu := gtk.NewMenu()
    cascademenu.SetSubmenu(submenu)

    menuitem = gtk.NewMenuItemWithMnemonic("Q_uitter")
    menuitem.Connect("activate", func() {
        gtk.MainQuit()
    })
    submenu.Append(menuitem)


    cascademenu = gtk.NewMenuItemWithMnemonic("_Aide")
    menubar.Append(cascademenu)
    submenu = gtk.NewMenu()
    cascademenu.SetSubmenu(submenu)

    auteurs := gtk.NewEntry()
    auteurs.SetText("Application crée en MCS par Olivier CANO et Adrien CHAPELET")

    menuitem = gtk.NewMenuItemWithMnemonic("À_ propos")
    menuitem.Connect("activate", func() {
        messagedialog := gtk.NewMessageDialog(
            button.GetTopLevelAsWindow(),
            gtk.DIALOG_MODAL,
            gtk.MESSAGE_INFO,
            gtk.BUTTONS_OK,
            auteurs.GetText())
        messagedialog.Response(func() {
            messagedialog.Destroy()
        })
        messagedialog.Run()
    })
    submenu.Append(menuitem)


    //--------------------------------------------------------
    // GtkStatusbar
    //--------------------------------------------------------
    statusbar := gtk.NewStatusbar()
    context_id := statusbar.GetContextId("go-gtk")
    statusbar.Push(context_id, "En attente de synchronisation")

    framebox2.PackStart(statusbar, false, false, 0)

    //--------------------------------------------------------
    // Event
    //--------------------------------------------------------
    window.Add(vbox)
    window.SetSizeRequest(500, 300)
    window.ShowAll()
    gtk.Main()
}