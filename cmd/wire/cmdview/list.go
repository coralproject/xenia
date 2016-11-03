package cmdview

import (
	"encoding/json"

	"github.com/coralproject/shelf/internal/wire/view"
	"github.com/spf13/cobra"
)

var listLong = `Use list to list all the views avaiable.

Example:
	view list
`

// addList handles the list of views.
func addList() {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the views available.",
		Long:  listLong,
		Run:   runList,
	}

	viewCmd.AddCommand(cmd)
}

// runList is the code that implements the list command.
func runList(cmd *cobra.Command, args []string) {
	cmd.Printf("Listing All Views Available.\n")

	// Retrieves the current views from Mongo.
	results, err := view.GetAll("", mgoDB)
	if err != nil {
		cmd.Println("Listing Views : ", err)
		return
	}

	// Prepare the results for printing.
	data, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		cmd.Println("Listing Views : ", err)
		return
	}

	cmd.Printf("\n%s\n\n", string(data))
	cmd.Println("\n", "Listing Views : Listed")
}
