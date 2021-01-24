package cmd

import (
	"baoctl/pkg/config"
	"baoctl/pkg/tmpl"
	"baoctl/pkg/types"
	"baoctl/pkg/util"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func Run() error {
	return (&cli.App{
		Name:  "baoctl",
		Usage: "baoctl淘宝工具",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "config", Value: "config.yml"},
		},
		Before: func(ctx *cli.Context) error {
			return config.Initialize(&types.Options{
				FilePath: ctx.String("config"),
			})
		},
		Action: func(ctx *cli.Context) error {
			return Main(ctx)
		},
	}).Run(os.Args)
}

func Main(ctx *cli.Context) error {
	Instance().PrintCommands()
	for {
		cmd := 0
		fmt.Print(tmpl.CmdWait)
		if _, err := fmt.Scan(&cmd); err != nil {
			util.PrintError(err)
			continue
		}

		if err := Instance().Exec(ctx.Context, cmd); err != nil {
			fmt.Println()
			util.PrintError(err)
			continue
		}
		return nil
	}
}
