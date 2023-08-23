package start

import (
	"context"

	"github.com/solodba/mysql_install/protocol"
	"github.com/spf13/cobra"
)

var (
	ctx = context.Background()
)

// MySQL服务结构体
type Server struct {
	MySQLInstallSvc *protocol.MySQLInstallSvc
}

// MySQL服务结构体初始化函数
func NewServer() *Server {
	return &Server{
		MySQLInstallSvc: protocol.NewMySQLInstallSvc(),
	}
}

// 项目启动子命令
var Cmd = &cobra.Command{
	Use:     "start",
	Short:   "mysql_install start service",
	Long:    "mysql_install start service",
	Example: "mysql_install start -f etc/config.toml",
	RunE: func(cmd *cobra.Command, args []string) error {
		svc := NewServer()
		err := svc.MySQLInstallSvc.MySQLInstall(ctx)
		if err != nil {
			return err
		}
		return nil
	},
}
