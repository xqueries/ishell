package ishell

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// App describes the execution of IShell.
// This runs one command at a time. The commands
// can be chained together using context-chaining
// for all use-casses.
type App struct {
	Context
}

// Context describes a single command and supporting
// data that is needed to run that command.
type Context struct {
	Input   interface{}
	Command *cobra.Command
}

// IShell implements ISheller, an interactive shell.
// This takes in all functions to run and the commands
// at the time of creation.
type IShell struct {
	counter int
}

var _ ISheller = (*IShell)(nil)

// NewIShell returns a new instance of IShell.
func NewIShell() *IShell {
	return &IShell{
		counter: 1,
	}
}

// Start starts the IShell which includes the data to
// maintain a shell running with knowledge of all commands
// and functions to execute on receipt of all supported
// commands.
// It expects the argument of a slice of contexts, where
// the first context is the root context.
func (ish *IShell) Start(genCtx func() Context) error {

	var app *App
	app = &App{
		Context: Context{
			Input: promptui.Prompt{
				Label: "what do you want?",
			},
			Command: &cobra.Command{
				RunE: func(cmd *cobra.Command, args []string) error {
					fmt.Println("inside command, input:", args[0])
					if args[0] != "table" {
						return fmt.Errorf("we don't have '%s'", args[0])
					}
					app.Context = genCtx()
					return nil
				},
			},
		},
	}
	return app.Run()
}

// Run runs an instance of an IShell App.
func (app *App) Run() error {
	var err error
	var cmd *cobra.Command
	for {
		var result string
		switch prompt := app.Input.(type) {
		case promptui.Select:
			_, str, err := prompt.Run()
			// Close if it's an interrupt signal.
			if err == promptui.ErrInterrupt {
				return err
			}
			if err != nil {
				fmt.Println("error:", err)
				continue
			}
			result = str
		case promptui.Prompt:
			str, err := prompt.Run()
			// Close if it's an interrupt signal.
			if err == promptui.ErrInterrupt {
				return err
			}
			if err != nil {
				fmt.Println("error this:", err)
				continue
			}
			result = str
		}
		cmdNext := app.Context.Command
		if cmdNext != cmd {
			fmt.Println("command changed")
		}
		cmd = cmdNext

		cmd.SetArgs([]string{result})
		cmd.SilenceUsage = true
		err = cmd.Execute()
		if err != nil {
			break
		}
	}
	return err
}
