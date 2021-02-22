package parser

import (
	"fmt"
	"strings"

	"github.com/tucnak/climax"
)

var _ Parser = (*SimpleParser)(nil)

// SimpleParser uses "https://github.com/akamensky/argparse"
// to parse its input.
type SimpleParser struct {
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

func (sp *SimpleParser) Parse(input string) error {

	fmt.Println("reached parser")
	demo := climax.New(" ")
	demo.Brief = "Demo is a funky demonstation of Climax capabilities."
	demo.Version = "stable"

	joinCmd := climax.Command{
		Name:  "join",
		Brief: "merges the strings given",
		Usage: `[-s=] "a few" distinct strings`,
		Help:  `Lorem ipsum dolor sit amet amet sit todor...`,

		Flags: []climax.Flag{
			{
				Name:     "separator",
				Short:    "s",
				Usage:    `--separator="."`,
				Help:     `Put some separating string between all the strings given.`,
				Variable: true,
			},
		},

		Examples: []climax.Example{
			{
				Usecase:     `-s . "google" "com"`,
				Description: `Results in "google.com"`,
			},
		},

		Handle: func(ctx climax.Context) int {
			var separator string
			if sep, ok := ctx.Get("separator"); ok {
				separator = sep
			}

			fmt.Println(strings.Join(ctx.Args, separator))

			return 0
		},
	}

	demo.AddCommand(joinCmd)
	demo.Run()
	return nil
}

// // Create new parser object
// parser := argparse.NewParser("overview", "provides an overview of the xdb file")
// // Create string flag
// _ = parser.String("s", "string", &argparse.Options{Required: true, Help: "String to print"})
// // Parse input
// err := parser.Parse(os.Args)
// if err != nil {
// 	// In case of error print error and print usage
// 	// This can also be done by passing -h or --help flags
// 	fmt.Print(parser.Usage(err))
// }
// // Finally print the collected string
// // fmt.Println(*s)
