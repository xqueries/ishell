package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/xqueries/ishell/ishell"
)

var counter = 0
var (
	givenTables = []string{
		"table0",
	}
)

func dynamicCount() []string {
	counter++
	fmt.Println(counter)
	tbl := "table" + strconv.Itoa(counter)
	givenTables = append(givenTables, tbl)
	fmt.Println(givenTables)
	return givenTables
}

func main() {
	ish := ishell.NewIShell()

	// m := make(map[string]func() ishell.Context)

	err := ish.Start(
		func() ishell.Context {
			return generateContext("tableContext")
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func generateContext(label string) ishell.Context {
	var ctx ishell.Context
	switch label {
	case "context":
		fmt.Println("generating")
		var contextApp *ishell.App
		contextApp = &ishell.App{
			Context: ishell.Context{
				Input: promptui.Select{
					Label: "choose assssss",
					Items: dynamicCount(),
				},
				Command: &cobra.Command{
					RunE: func(cmd *cobra.Command, args []string) error {
						contextApxp.Context = generateContext("tableContext")
						return nil
					},
				},
			},
		}
		ctx = contextApp.Context
	case "tableContext":
		var tableContextApp *ishell.App
		tableContextApp = &ishell.App{
			Context: ishell.Context{
				Input: promptui.Select{
					Label: "choose a table",
					Items: dynamicCount(),
				},
				Command: &cobra.Command{
					RunE: func(cmd *cobra.Command, args []string) error {
						fmt.Println("inside table command, desired:", args[0])
						tableContextApp.Context = generateContext("context")
						return nil
					},
				},
			},
		}
		ctx = tableContextApp.Context
	}
	return ctx
}
