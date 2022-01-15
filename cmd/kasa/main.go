package main

import (
	"os"

	"github.com/alecthomas/kong"
	"github.com/posener/complete"
	"github.com/willabides/kongplete"
	"github.com/winebarrel/kasa"
	"github.com/winebarrel/kasa/esa"
	"github.com/winebarrel/kasa/subcmd"
)

var version string

var cli struct {
	Version            kong.VersionFlag
	Team               string                       `required:"" env:"ESA_TEAM" help:"esa team"`
	Token              string                       `required:"" env:"ESA_TOKEN" help:"esa access token"`
	Debug              bool                         `help:"Debug flag."`
	Append             subcmd.AppendCmd             `cmd:"" help:"Append text to a post."`
	Cat                subcmd.CatCmd                `cmd:"" help:"Print a post."`
	Comment            subcmd.CommentCmd            `cmd:"" help:"Comment to a post."`
	Cp                 subcmd.CpCmd                 `cmd:"" help:"Copy posts."`
	Info               subcmd.InfoCmd               `cmd:"" help:"Show a post info."`
	Ls                 subcmd.LsCmd                 `cmd:"" help:"List posts."`
	Mv                 subcmd.MvCmd                 `cmd:"" help:"Move posts."`
	Mvcat              subcmd.MvcatCmd              `cmd:"" help:"Rename a category."`
	Open               subcmd.OpenCmd               `cmd:"" help:"Open a post in the browser."`
	Post               subcmd.PostCmd               `cmd:"" help:"Create/Update a post."`
	Rm                 subcmd.RmCmd                 `cmd:"" help:"Delete posts."`
	Rmi                subcmd.RmiCmd                `cmd:"" help:"Delete a post by number."`
	Search             subcmd.SearchCmd             `cmd:"" help:"Search posts."`
	Tag                subcmd.TagCmd                `cmd:"" help:"Tagging posts."`
	InstallCompletions kongplete.InstallCompletions `cmd:"" help:"Install shell completions"`
}

func main() {
	parser := kong.Must(&cli, kong.Vars{"version": version})

	kongplete.Complete(parser,
		kongplete.WithPredictor("file", complete.PredictFiles("*")),
	)

	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	err = ctx.Run(&kasa.Context{
		Team:   cli.Team,
		Driver: esa.NewDriver(cli.Team, cli.Token, cli.Debug),
		Fmt:    &kasa.PrinterImpl{},
	})

	ctx.FatalIfErrorf(err)
}
