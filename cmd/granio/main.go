// Run a command line helper for various Fyne tools.
package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/urfave/cli/v2"

	"github.com/Uzer007-admin/granio-tools/cmd/granio/internal/commands"
)

func main() {
	verCmd := commands.Version()
	app := &cli.App{
		Name:        "granio",
		Usage:       "Сборка и упаковка приложений Granio.",
		Description: "Средства разработки и распространения приложений Granio.",
		Commands: []*cli.Command{
			commands.Init(),
			commands.Env(),
			commands.Build(),
			commands.Package(),
			commands.Release(),
			commands.Install(),
			commands.Serve(),
			commands.Translate(),
			verCmd,
			commands.Bundle(),
		},
	}

	info, ok := debug.ReadBuildInfo()
	if ok {
		app.Version = info.Main.Version
	}

	cli.VersionPrinter = func(ctx *cli.Context) {
		exitOnError(verCmd.Action(ctx))
	}

	exitOnError(app.Run(os.Args))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
