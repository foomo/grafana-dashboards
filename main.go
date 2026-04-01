package main

import (
	"fmt"
	"os"
	"runtime/debug"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
	"github.com/foomo/grafana-dashboards/cmd"
	pkgcmd "github.com/foomo/grafana-dashboards/internal/cmd"
	"github.com/pterm/pterm"
)

func init() {
	pkgcmd.Init()
}

func main() {
	root := cmd.NewRoot()
	root.AddCommand(
		cmd.NewConfig(),
		cmd.NewGenerate(),
		cmd.NewVersion(),
	)

	say := func(msg string) string {
		if say, cerr := cowsay.Say(msg, cowsay.BallonWidth(80)); cerr == nil {
			msg = say
		}

		return msg
	}

	code := 0

	defer func() {
		if r := recover(); r != nil {
			pterm.Error.Println(say("It's time to panic"))
			pterm.Error.Println(fmt.Sprintf("%v", r))
			pterm.Error.Println(string(debug.Stack()))

			code = 1
		}

		os.Exit(code)
	}()

	if err := root.Execute(); err != nil {
		pterm.Error.Println(say("Ups, something went wrong"))
		pterm.Error.Println(err.Error())

		code = 1
	}
}
