package start

import (
	"github.com/spf13/cobra"
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "mysql_install start service",
	Long:    "mysql_install start service",
	Example: "mysql_install start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
