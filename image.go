package main

import (
	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/canvas"
)

// Declare conformity with editor interface
var _ editor = (*imageEditor)(nil)

type imageEditor struct {
	uri   granio.URI
	image *canvas.Image
}

func newImageEditor(u granio.URI, _ granio.Window) editor {
	img := canvas.NewImageFromURI(u)
	img.FillMode = canvas.ImageFillContain
	editor := &imageEditor{uri: u, image: img}

	return editor
}

func (i *imageEditor) changed() bool {
	return false
}

func (i *imageEditor) content() granio.CanvasObject {
	return i.image
}

func (i *imageEditor) close() {
}

func (i *imageEditor) run() {
}

func (i *imageEditor) save() {
}
