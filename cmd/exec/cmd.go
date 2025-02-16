package exec

import (
	"slices"

	itbasisCoreCmd "github.com/itbasis/go-tools-core/cmd"
	itbasisCoreOs "github.com/itbasis/go-tools-core/os"
	sdkmCmd "github.com/itbasis/go-tools-sdkm/internal/cmd"
	"github.com/itbasis/go-tools-sdkm/plugins"
	"github.com/spf13/cobra"
)

func NewExecCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "exec",
		Short: "Execute a command in a plugin",
	}

	sdkmCmd.InitFlagRebuildCache(cmd.PersistentFlags())

	plugins.AddPluginsAsSubCommands(
		cmd, func(cmdChild *cobra.Command) {
			cmdChild.DisableFlagParsing = true
			cmdChild.Use = itbasisCoreCmd.BuildUse(cmdChild.Use, "{<program>}", "[<args...>]")
			cmdChild.Args = cobra.MinimumNArgs(1)
			cmdChild.ArgAliases = []string{"program"}
			cmdChild.RunE = _runE
		},
	)

	return cmd
}

func _runE(cmd *cobra.Command, args []string) error {
	argDebugFlagName := "--" + sdkmCmd.FlagDebugName

	//nolint:wrapcheck // TODO
	return plugins.GetPluginByID(cmd).
		Exec(
			cmd.Context(),
			sdkmCmd.IsFlagRebuildCache(cmd),
			itbasisCoreOs.Pwd(),
			cmd.InOrStdin(),
			cmd.OutOrStdout(),
			cmd.OutOrStderr(),
			slices.DeleteFunc(
				args, func(s string) bool {
					return s == argDebugFlagName
				},
			),
		)
}
