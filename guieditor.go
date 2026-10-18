package main

import (
	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/dialog"

	"github.com/Uzer007-admin/granio-studio/internal/guibuilder"
)

// Declare conformity with editor interface
var _ editor = (*guiEditor)(nil)

type guiEditor struct {
	uri     granio.URI
	builder *guibuilder.Builder
	edited  bool
	win     granio.Window
}

func newGuiEditor(u granio.URI, win granio.Window) editor {
	builder := guibuilder.NewBuilder(u, win)
	return &guiEditor{uri: u, builder: builder, win: win}
}

func (g *guiEditor) changed() bool {
	return g.edited
}

func (g *guiEditor) content() granio.CanvasObject {
	return g.builder.MakeUI()
}

func (g *guiEditor) close() {
}

func (g *guiEditor) run() {
	g.builder.Run()
}

func (g *guiEditor) save() {
	err := g.builder.Save()
	if err != nil {
		dialog.ShowError(err, g.win)
		return
	}

	g.edited = false
}
