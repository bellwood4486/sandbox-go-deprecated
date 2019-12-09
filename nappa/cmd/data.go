package cmd

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Generates data",
	Long:  `This command generates data for request.`,
	Run: func(cmd *cobra.Command, args []string) {
		out(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func out(w io.Writer) {
	signal.Ignore(syscall.SIGPIPE)

	for i := 0; i < 5; i++ {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				break
			}
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
