package mysql

import "context"

const (
	AppName = "mysql"
)

type Service interface {
	// 解压MySQL压缩文件
	UnzipMySQLFile(ctx context.Context) error
	// 创建MySQL相关目录
	CreateMySQLDir(ctx context.Context) error
	// 判断是否有MySQL进程
	IsMySQLRun(ctx context.Context) error
	// 创建MySQL用户
	CreateMySQLUser(ctx context.Context) error
	// 修改权限
	ChangeMySQLDirPerm(ctx context.Context) error
	// MySQL初始化
	InitialMySQL(ctx context.Context) error
	// 启动MySQL
	StartMySQL(ctx context.Context) error
}
