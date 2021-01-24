package command

import (
	"context"
	"fmt"
)

type Command struct {
	Code   int
	Desc   string
	Action func(ctx context.Context, args ...interface{}) error
}

func (c *Command) String() string {
	return fmt.Sprintf("[%d]%s", c.Code, c.Desc)
}
