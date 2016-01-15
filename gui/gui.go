// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

// Create a PanelHolder with a 3 panels
func panelHolder(name string, theme gxui.Theme) gxui.PanelHolder {
	label := func(text string) gxui.Label {
		label := theme.CreateLabel()
		label.SetText(text)
		return label
	}

	holder := theme.CreatePanelHolder()
	holder.AddPanel(label(name+" Tout"), name+" Tous mes fichiers")
	holder.AddPanel(label(name+" Cours"), name+" Mes cours")
	holder.AddPanel(label(name+" Musique"), name+" Ma musique")
	holder.AddPanel(label(name+" Photos"), name+" Mes photos")
	holder.AddPanel(label(name+" Vidéos"), name+" Mes vidéos")
	return holder
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	label := theme.CreateLabel()
	label.SetText("Clou")

	splitterAB := theme.CreateSplitterLayout()
	splitterAB.SetOrientation(gxui.Horizontal)
	splitterAB.AddChild(panelHolder("Local", theme))
	splitterAB.AddChild(panelHolder("Cloud", theme))

	vSplitter := theme.CreateSplitterLayout()
	vSplitter.SetOrientation(gxui.Vertical)
	vSplitter.AddChild(splitterAB)

	window := theme.CreateWindow(800, 600, "Panels")
	window.AddChild(label)
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(vSplitter)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}
