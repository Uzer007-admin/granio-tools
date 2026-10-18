package main

import (
	"strings"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/storage"
	"github.com/Uzer007-admin/granio/v2/widget"
)

const preferenceKeyRecent = "recentlist"

func addRecent(u granio.URI, p granio.Preferences) {
	old := loadRecents(p)
	if len(old) > 0 {
		if old[0].String() == u.String() {
			return
		}

		for i, u2 := range old {
			if u2.String() == u.String() {
				if i < len(old)-1 {
					old = old[:i]
				} else {
					old = append(old[:i], old[i+1:]...)
				}
			}
		}
	}

	all := append([]granio.URI{u}, old...)
	str := ""
	for _, u := range all {
		str += "|" + u.String()
	}
	p.SetString(preferenceKeyRecent, str[1:])
}

func loadRecents(p granio.Preferences) []granio.URI {
	var ret []granio.URI
	val := p.String(preferenceKeyRecent)
	for _, s := range strings.Split(val, "|") {
		u, err := storage.ParseURI(s)
		if err == nil {
			ret = append(ret, u)
		}
	}
	return ret
}

func makeRecentList(p granio.Preferences, f func(granio.URI)) *widget.List {
	recents := loadRecents(p)
	r := widget.NewList(
		func() int {
			return len(recents)
		},
		func() granio.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, o granio.CanvasObject) {
			name := recents[id].Name()
			o.(*widget.Label).SetText(name)
		})
	r.OnSelected = func(id widget.ListItemID) {
		f(recents[id])
	}
	return r
}
