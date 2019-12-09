package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// vreqCmd represents the vreq command
var vreqCmd = &cobra.Command{
	Use:   "vreq",
	Short: "Generates a request for vegeta",
	Long:  `This command generates a vegeta request.`,
	Run: func(cmd *cobra.Command, args []string) {
		in(os.Stdin)
	},
}

func init() {
	rootCmd.AddCommand(vreqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vreqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vreqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func in(r io.Reader) {
	for {
		var s string
		_, err := fmt.Fscanln(r, &s)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("readed: %s\n", s)
	}
}
