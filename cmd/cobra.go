package cmd

import (
	"errors"

	"github.com/haobinfei/gin-test/pkg/logger"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "gin-test",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `gin-test`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用 ferry, 可以使用 -h 查看命令`
		logger.Infof("%s/n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand()
}

func Execute() {

}
