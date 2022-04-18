package runners

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/aawadallak/simple-cli-tool/services"
)

type consumer struct {
	fs *flag.FlagSet

	action    string
	topic     string
	qntityMsg int
	arn       string
}

func NewConsumer() Runner {
	flag := &consumer{
		fs: flag.NewFlagSet("consume", flag.ContinueOnError),
	}

	flag.fs.StringVar(&flag.action, "action", "", "is used to set action (e.g single, batch)")
	flag.fs.StringVar(&flag.topic, "topic", "", "")
	flag.fs.IntVar(&flag.qntityMsg, "quantity", 10, "set the quantity message retrieved by the consumer, by default is 10")
	flag.fs.StringVar(&flag.arn, "msg", "", "")

	return flag
}

func (g *consumer) Name() string {
	return g.fs.Name()
}

func (g *consumer) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *consumer) Run() error {
	if err := g.validate(); err != nil {
		return err
	}

	do, ok := services.ConsumerAction[g.action]
	if !ok {
		return fmt.Errorf("must provide a valid action")
	}

	return do(context.Background())
}

func (g *consumer) validate() error {
	if g.topic != "" && g.qntityMsg == 0 {
		return errors.New("you must provide quantity of messages to be retrevied")
	}
	return nil
}
