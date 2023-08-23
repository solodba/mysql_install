package protocol

import (
	"context"

	"github.com/solodba/mcube/apps"
	"github.com/solodba/mysql_install/apps/mysql"
)

// MySQL安装服务结构体
type MySQLInstallSvc struct {
	svc mysql.Service
}

// MySQL安装服务结构体构造函数
func NewMySQLInstallSvc() *MySQLInstallSvc {
	return &MySQLInstallSvc{
		svc: apps.GetInternalApp(mysql.AppName).(mysql.Service),
	}
}

// MySQL安装服务
func (m *MySQLInstallSvc) MySQLInstall(ctx context.Context) error {
	err := m.svc.UnzipMySQLFile(ctx)
	if err != nil {
		return err
	}
	err = m.svc.CreateMySQLDir(ctx)
	if err != nil {
		return err
	}
	err = m.svc.IsMySQLRun(ctx)
	if err != nil {
		return err
	}
	err = m.svc.CreateMySQLUser(ctx)
	if err != nil {
		return err
	}
	err = m.svc.ChangeMySQLDirPerm(ctx)
	if err != nil {
		return err
	}
	err = m.svc.InitialMySQL(ctx)
	if err != nil {
		return err
	}
	err = m.svc.StartMySQL(ctx)
	if err != nil {
		return err
	}
	return nil
}
