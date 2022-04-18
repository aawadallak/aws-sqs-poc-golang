package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/aawadallak/simple-cli-tool/runners"
)

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func root(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a sub-command")
	}

	runners := []runners.Runner{
		runners.NewCreate(),
		runners.NewLister(),
		runners.NewConsumer(),
		runners.NewPublisher(),
	}

	subcommand := os.Args[1]

	for _, cmd := range runners {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("unknown subcommand: %s", subcommand)
}
