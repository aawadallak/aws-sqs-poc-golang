package runners

import (
	"context"
	"flag"
	"fmt"

	"github.com/aawadallak/simple-cli-tool/services"
)

type publisher struct {
	fs *flag.FlagSet

	action string
	topic  string
	arn    string
}

func NewPublisher() Runner {
	flag := &publisher{
		fs: flag.NewFlagSet("publish", flag.ExitOnError),
	}

	flag.fs.StringVar(&flag.action, "action", "", "is used to set action (e.g single, batch)")
	flag.fs.StringVar(&flag.arn, "msg", "", "")

	return flag
}

func (g *publisher) Name() string {
	return g.fs.Name()
}

func (g *publisher) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *publisher) Run() error {
	do, ok := services.PublisherAction[g.action]
	if !ok {
		return fmt.Errorf("must provide a valid action")
	}

	return do(context.Background())
}
