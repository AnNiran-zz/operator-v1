package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var err error

var rootCmd = &cobra.Command{
	Use: "",
}

var crdCmd = &cobra.Command{
	Use:   "crd",
	Short: "Manage Custom Resource Definitions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// args
		// [action] [crd]
		if len(args) != 2 {
			printHelpCRD()
			os.Exit(1)
		}

		action := args[0]
		switch action {
		case "add":
			err = addCRD(args[1])

		case "remove":
			err = removeCRD(args[1])

		case "generate-client":
			err = generateCRDClientData(args[1])

		case "delete-client":
			err = deleteCRDClientData(args[1])

		default:
			fmt.Println("Unknown command")
			printHelpCRD()
			os.Exit(1)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "Manage Custom Objects",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// args
		// [action] [crd] [object]
		if len(args) != 3 {
			printHelpObject()
			os.Exit(1)
		}

		action := args[0]
		switch action {
		case "add":
			err = createObject(args[1:])

		case "remove":
			err = deleteObject(args[1:])

		case "update":
			err = updateObject(args[1:])

		default:
			fmt.Println("Unknown command")
			printHelpObject()
			os.Exit(1)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Run starts CLI
func Run() {
	rootCmd.AddCommand(crdCmd)
	rootCmd.AddCommand(objectCmd)
	rootCmd.Execute()
}
