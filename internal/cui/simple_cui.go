package cui

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/xqueries/ishell/internal/parser"
)

var _ Cui = (*SimpleCui)(nil)

type SimpleCui struct {
}

func NewSimpleCui() *SimpleCui {
	return &SimpleCui{}
}

func (sc *SimpleCui) Start() error {
	validate := func(input string) error {
		// _, err := strconv.ParseFloat(input, 64)
		// if err != nil {
		// 	return errors.New("Invalid number")
		// }
		return nil
	}

	for {
		prompt := promptui.Prompt{
			Label:    "Number",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		parser := parser.NewSimpleParser()
		err = parser.Parse(result)
	}
}
