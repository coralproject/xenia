package cmdquery

import "github.com/spf13/cobra"

// envKey defines the environment variable to be looked for, to load rules
// from if provided.
var envKey = "SHELF_PATH"

// defDir defines the default directory to search for.
var defDir = "queries"

// queryCmd represents the parent for all query cli commands.
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query provides a shelf CLI for managing and executing queries.",
}

// GetCommands returns the query commands.
func GetCommands() *cobra.Command {
	addCreate()
	addGet()
	addDel()
	addExecute()
	addUpd()
	addList()
	return queryCmd
}
