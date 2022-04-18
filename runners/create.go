package runners

import "flag"

type create struct {
	fs *flag.FlagSet

	topic string
	arn   string
}

func NewCreate() Runner {
	flag := &create{
		fs: flag.NewFlagSet("create", flag.ContinueOnError),
	}

	flag.fs.StringVar(&flag.topic, "topic", "", "")
	flag.fs.StringVar(&flag.arn, "msg", "", "")

	return flag
}

func (g *create) Name() string {
	return g.fs.Name()
}

func (g *create) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *create) Run() error {
	return nil
}
