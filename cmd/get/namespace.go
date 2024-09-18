package get

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns"},
	Short:   "Short namespace",
	Long: `Long
	namespace`,
	Run: func(cmd *cobra.Command, args []string) {
		dirFile, err := NewSSHFileInfo()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(*dirFile)
	},
}

func init() {
	getCmd.AddCommand(namespaceCmd)
}
