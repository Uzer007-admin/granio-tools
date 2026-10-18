package main

import "github.com/Uzer007-admin/granio/v2"

var editorsByFilename = map[string]func(granio.URI, granio.Window) editor{
	".gui.json": newGuiEditor,
	"go.mod":    newTextEditor,
}

var editorsByMime = map[string]func(granio.URI, granio.Window) editor{
	"application/xml-tdd": newTextEditor, // Go mime on older systems?
	"image/jpeg":          newImageEditor,
	"image/png":           newImageEditor,
	"image/svg+xml":       newImageEditor,
	"text/markdown":       newTextEditor,
	"text/plain":          newTextEditor,
	"text/x-go":           newTextEditor,
}

type editor interface {
	changed() bool
	close()
	content() granio.CanvasObject
	run()
	save()
}
