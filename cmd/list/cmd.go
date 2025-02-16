package list

import (
	itbasisCoreCmd "github.com/itbasis/go-tools-core/cmd"
	sdkmCmd "github.com/itbasis/go-tools-sdkm/internal/cmd"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "list",
		Short:  "List installed versions",
		PreRun: itbasisCoreCmd.LogCommand,
	}

	sdkmCmd.InitFlagRebuildCache(cmd.PersistentFlags())

	cmd.AddCommand(newListAllCommand())

	return cmd
}
