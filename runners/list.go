package runners

import "flag"

type list struct {
	fs *flag.FlagSet

	topic string
	arn   string
}

func NewLister() Runner {
	flag := &list{
		fs: flag.NewFlagSet("list", flag.ExitOnError),
	}

	flag.fs.StringVar(&flag.topic, "topic", "", "")
	flag.fs.StringVar(&flag.arn, "msg", "", "")

	return flag
}

func (g *list) Name() string {
	return g.fs.Name()
}

func (g *list) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *list) Run() error {
	return nil
}
