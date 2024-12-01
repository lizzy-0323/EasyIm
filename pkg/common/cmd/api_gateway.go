package cmd

import (
	"github.com/spf13/cobra"
)

func NewApiGateWay() *cobra.Command {
	return &cobra.Command{
		Use:   "api-gateway",
		Short: "Start the api gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
