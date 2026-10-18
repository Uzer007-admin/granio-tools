package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/mod/modfile"

	"github.com/Uzer007-admin/granio-tools/cmd/granio/internal/metadata"
	"github.com/Uzer007-admin/granio-tools/cmd/granio/internal/templates"
	"github.com/Uzer007-admin/granio-tools/cmd/granio/internal/util"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "Создаёт новое приложение Granio",
		ArgsUsage: "[module-path]",
		Action:    initAction,
		Description: "Создаёт приложение Granio в текущем каталоге: go.mod, main.go и FyneApp.toml.\n" +
			"Существующие файлы сохраняются.",
		Flags: []cli.Flag{
			stringFlags["app-id"](nil),
			stringFlags["name"](nil),
			stringFlags["icon"](nil),
			boolFlags["verbose"](nil),
		},
	}
}

func getAppID(modpath string) string {
	p := util.SplitSlash(strings.ReplaceAll(modpath, "-", "_"))
	if len(p) == 0 {
		return ""
	}

	d := util.SplitDot(p[0])
	r := make([]string, len(p)+len(d)-1)
	for n, e := range d {
		r[len(d)-n-1] = e
	}
	for n, e := range p {
		if n == 0 {
			continue
		}
		r[len(d)+n-1] = e
	}

	return strings.Join(r, ".")
}

func getAppName(modpath string) string {
	p := util.SplitSlash(modpath)
	if len(p) == 0 {
		return ""
	}

	if len(p) > 1 {
		return p[len(p)-1]
	}

	d := util.SplitDot(p[0])

	return d[0]
}

func checkFileOrDo(file string, cb func() error) error {
	fi, err := os.Stat(file)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if fi != nil {
		return nil
	}

	return cb()
}

func initAction(ctx *cli.Context) error {
	modpath := ctx.Args().Get(0)
	appID := ctx.String("app-id")
	appName := ctx.String("name")
	icon := ctx.String("icon")
	verbose := ctx.Bool("verbose")

	if modpath == "" {
		modpath = "example"

		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		if wd != "" {
			modpath = filepath.Base(wd)
		}
	}

	if appID == "" {
		appID = getAppID(modpath)
	}

	if appName == "" {
		appName = getAppName(modpath)
	}

	data := &metadata.FyneApp{}
	data.Details.ID = appID
	data.Details.Name = appName
	data.Details.Icon = icon
	data.Details.Version = "0.0.0"
	data.Migrations = map[string]bool{
		"fyneDo": true,
	}

	if err := checkFileOrDo("main.go", func() error {
		f, err := os.Create("main.go")
		if err != nil {
			return err
		}
		defer f.Close()
		return templates.HelloWorld.Execute(f, &data)
	}); err != nil {
		return err
	}

	if err := checkFileOrDo("go.mod", func() error {
		cmd := exec.Command("go", "mod", "init", modpath)
		if verbose {
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		}

		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("go", "mod", "edit", "-go=1.27.1")
		if verbose {
			cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
		}
		return cmd.Run()
	}); err != nil {
		return err
	}

	if err := ensureGranioDependency(".", verbose); err != nil {
		return err
	}

	if err := checkFileOrDo("FyneApp.toml", func() error {
		return metadata.SaveStandard(data, ".")
	}); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "tidy")
	if verbose {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	if err := os.Mkdir("translations", util.DirPermDefault); err != nil {
		return err
	}

	args := []string{ctx.App.Name, "translate"}
	if verbose {
		args = append(args, "-v")
	}
	args = append(args, "translations/en.json")
	args = append(args, "main.go")
	if err := ctx.App.Run(args); err != nil {
		return fmt.Errorf("failed to run command: %v", err)
	}

	fmt.Println("Приложение готово. Запуск: go run .")

	return nil
}

// granioVersionForInit avoids selecting an inherited upstream release through @latest.
func granioVersionForInit() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, dep := range info.Deps {
			if dep.Path == fyneModule && strings.HasPrefix(dep.Version, "v2.") {
				return dep.Version
			}
		}
	}
	return "v2.8.2-0.20261018000600-86cc66cfefa1"
}

// ensureGranioDependency preserves a project-selected version and pins the CLI's
// core version when the application does not yet declare one.
func ensureGranioDependency(directory string, verbose bool) error {
	data, err := os.ReadFile(filepath.Join(directory, "go.mod"))
	if err != nil {
		return err
	}
	mod, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		return err
	}
	for _, dep := range mod.Require {
		if dep.Mod.Path == fyneModule {
			return nil
		}
	}
	cmd := exec.Command("go", "mod", "edit", "-require="+fyneModule+"@"+granioVersionForInit())
	cmd.Dir = directory
	if verbose {
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	}
	return cmd.Run()
}
