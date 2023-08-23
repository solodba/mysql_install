package mysql

import "context"

const (
	AppName = "mysql"
)

type Service interface {
	// 解压MySQL压缩文件
	UnzipMySQLFile(context.Context) error
	// 创建MySQL相关目录
	CreateMySQLDir(context.Context) error
	// 判断是否有MySQL进程
	IsMySQLRun(context.Context) error
	// 创建MySQL用户
	CreateMySQLUser(context.Context) error
	// 修改权限
	ChangeMySQLDirPerm(context.Context) error
	// MySQL初始化
	InitialMySQL(context.Context) error
	// 启动MySQL
	StartMySQL(context.Context) error
	// 增加环境量变量
	AddEnv(context.Context) error
}
