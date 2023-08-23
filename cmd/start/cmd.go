package start

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/mysql_install/apps/mysql"
	"github.com/spf13/cobra"
)

var (
	svc = apps.GetInternalApp(mysql.AppName).(mysql.Service)
	ctx = context.Background()
)

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "mysql_install start service",
	Long:    "mysql_install start service",
	Example: "mysql_install start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := svc.UnzipMySQLFile(ctx)
		if err != nil {
			return err
		}
		err = svc.CreateMySQLDir(ctx)
		if err != nil {
			return err
		}
		err = svc.IsMySQLRun(ctx)
		if err != nil {
			return err
		}
		err = svc.CreateMySQLUser(ctx)
		if err != nil {
			return err
		}
		err = svc.ChangeMySQLDirPerm(ctx)
		if err != nil {
			return err
		}
		err = svc.InitialMySQL(ctx)
		if err != nil {
			return err
		}
		return nil
	},
}
