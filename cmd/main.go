package main

import (
	"fmt"
	"os"

	arg "github.com/alexflint/go-arg"
	"github.com/vitorfox/gobadge/internal/badges"
	"github.com/vitorfox/gobadge/package/svg"
)

type Args struct {
	Config string `arg:"required"`
}

func main() {

	args := &Args{}
	arg.MustParse(args)

	conf := getConfig(args.Config)

	for _, b := range conf {

		var logic badges.Logicer

		switch b.Logic {
		case "moreWorst":
			logic = badges.NewLogicMoreWorstFromValues(b.Values)
		default:
			panic(fmt.Sprintf("Invalid logic:%s\n", b.Logic))
		}

		var sa *svg.Node

		switch b.Type {
		case "default":
			s := badges.NewDefault(b.Name)
			sa = s.Build(logic.GetParams())
		default:
			panic(fmt.Sprintf("Invalid type:%s\n", b.Type))
		}

		f, err := os.Create(b.Output)

		if err != nil {
			panic(err)
		}

		sa.Build(f)
	}
}
