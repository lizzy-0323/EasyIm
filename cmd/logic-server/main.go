package main

import (
	"go-im/internal/logic"

	"github.com/spf13/cobra"
)

// start logic server
func main() {
	logicCmd := NewLogicServerCmd()
	err := logicCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewLogicServerCmd() *cobra.Command {
	logicServerCmd := &cobra.Command{
		Use:   "logic-server",
		Short: "Start the logic server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := logic.Start()
			return err
		},
	}
	return logicServerCmd
}
