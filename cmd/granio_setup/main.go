package main

import (
	"github.com/Uzer007-admin/granio/v2/app"

	"github.com/Uzer007-admin/granio-tools/setup/pkg"
)

func main() {
	a := app.NewWithID("io.fyne.setup")
	pkg.ShowSummaryWindow(a)
	a.Run()
}
