package main

import (
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/dialog"
	"github.com/Uzer007-admin/granio/v2/storage"
	"github.com/Uzer007-admin/granio/v2/theme"
	"github.com/Uzer007-admin/granio/v2/widget"

	setup "github.com/Uzer007-admin/granio-setup/pkg"
)

func (d *defyne) menuActionNew() {
	input := widget.NewEntry()
	typeNames := make([]string, len(templates))
	for i, t := range templates {
		typeNames[i] = t.name
	}

	types := widget.NewSelect(typeNames, func(s string) {
		name := strings.Split(input.Text, ".")[0]
		for _, t := range templates {
			if t.name == s {
				name += t.ext
				continue
			}
		}
		input.SetText(name)
	})
	dialog.ShowForm("New file name", "Create", "Cancel",
		[]*widget.FormItem{
			widget.NewFormItem("File type", types),
			widget.NewFormItem("File name", input),
		},
		func(ok bool) {
			if !ok || input.Text == "" {
				return
			}

			uri, err := storage.Child(d.projectRoot, input.Text)
			if err != nil {
				dialog.ShowError(err, d.win)
				return
			}

			w, err := storage.Writer(uri)
			if err != nil {
				dialog.ShowError(err, d.win)
				return
			}
			_, _ = w.Write([]byte{})
			_ = w.Close()

			d.openEditor(uri)
			d.fileTree.Refresh()
		}, d.win)
}

func (d *defyne) menuActionRunProject() {
	pwd, _ := os.Getwd()
	os.Chdir(d.projectRoot.Path())
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()

	cmd = exec.Command("go", "run", ".")
	cmd.Dir = d.projectRoot.Path()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Start()

	os.Chdir(pwd)
}

func (d *defyne) menuActionRun() {
	if ed, ok := d.openEditors[d.fileTabs.Selected()]; ok {
		ed.run()
	}
}

func (d *defyne) menuActionSave() {
	if ed, ok := d.openEditors[d.fileTabs.Selected()]; ok {
		ed.save()
	}
}

func (d *defyne) menuActionFullScreenToggle() {
	d.win.SetFullScreen(!d.win.FullScreen())
}

func (d *defyne) makeHelpMenu(w granio.Window) *granio.Menu {
	return granio.NewMenu("Help",
		granio.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://fyne.io/defyne")
			_ = granio.CurrentApp().OpenURL(u)
		}),
		granio.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://fyne.io/support")
			_ = granio.CurrentApp().OpenURL(u)
		}),
		granio.NewMenuItemSeparator(),
		granio.NewMenuItem("Check Environment...", func() {
			if w != nil {
				setup.ShowSummaryDialog(w)
			} else {
				setup.ShowSummaryWindow(granio.CurrentApp())
			}
		}),
		granio.NewMenuItemSeparator(),
		granio.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://fyne.io/sponsor")
			_ = granio.CurrentApp().OpenURL(u)
		}),
	)
}

func (d *defyne) makeMenu() *granio.MainMenu {
	menu := granio.NewMainMenu(
		granio.NewMenu("File",
			granio.NewMenuItem("Open Project...", d.showProjectSelect),
			granio.NewMenuItemSeparator(),
			granio.NewMenuItem("New File...", d.menuActionNew),
			granio.NewMenuItemSeparator(),
			granio.NewMenuItem("Save", d.menuActionSave),
			granio.NewMenuItemSeparator(),
			granio.NewMenuItem("Run", d.menuActionRun),
			granio.NewMenuItem("Run Project", d.menuActionRunProject),
		))
	if runtime.GOOS != "darwin" {
		menu.Items = append(menu.Items,
			granio.NewMenu("Window",
				granio.NewMenuItem("Full Screen", d.menuActionFullScreenToggle),
			))
	}
	menu.Items = append(menu.Items,
		d.makeHelpMenu(d.win),
	)
	return menu
}

func (d *defyne) makeToolbar() *widget.Toolbar {
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), d.menuActionNew),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), d.menuActionSave),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), d.menuActionRun),
		widget.NewToolbarAction(theme.NewThemedResource(resourceFolderPlaySvg), d.menuActionRunProject))
}
