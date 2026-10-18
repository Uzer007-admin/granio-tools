package main

import (
	"strings"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/container"
	"github.com/Uzer007-admin/granio/v2/dialog"
	"github.com/Uzer007-admin/granio/v2/theme"

	xWidget "github.com/Uzer007-admin/granio-x/widget"
)

type fileTab struct {
	editor
	uri granio.URI
}

type defyne struct {
	win         granio.Window
	projectRoot granio.URI
	fileTabs    *container.DocTabs
	fileTree    *xWidget.FileTree
	openEditors map[*container.TabItem]*fileTab
}

func (d *defyne) openEditor(u granio.URI) {
	for tab, item := range d.openEditors {
		if item.uri.String() == u.String() {
			d.fileTabs.Select(tab)
			return
		}
	}

	var ed editor
	for name, e := range editorsByFilename {
		if strings.Contains(u.Name(), name) {
			ed = e(u, d.win)
		}
	}
	if ed == nil {
		if _, ok := editorsByMime[u.MimeType()]; !ok {
			dialog.ShowInformation("No registered editor",
				"No known editor for mime "+u.MimeType(), d.win)
			return
		}

		ed = editorsByMime[u.MimeType()](u, d.win)
	}

	newTab := container.NewTabItemWithIcon(u.Name(), theme.FileTextIcon(), ed.content())
	d.openEditors[newTab] = &fileTab{ed, u}

	d.fileTabs.Append(newTab)
	d.fileTabs.Select(newTab)
}
