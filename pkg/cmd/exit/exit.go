package exit

import (
	"baoctl/pkg/cmd"
	"baoctl/pkg/cmd/command"
	"context"
	"os"
)

func init() {
	cmd.RegisterCommand(&command.Command{
		Code: 99,
		Desc: `退出工具`,
		Action: func(ctx context.Context, args ...interface{}) error {
			os.Exit(0)
			return nil
		},
	})
}
