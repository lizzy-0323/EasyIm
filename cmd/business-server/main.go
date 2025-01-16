package main

import (
	"go-im/internal/business"

	"github.com/spf13/cobra"
)

func main() {
	businessCmd := NewBusinessServer()
	err := businessCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewBusinessServer() *cobra.Command {
	businessServerCmd := &cobra.Command{
		Use:   "business-server",
		Short: "Start the logic server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := business.Start(cmd.Context())
			return err
		},
	}
	return businessServerCmd
}
