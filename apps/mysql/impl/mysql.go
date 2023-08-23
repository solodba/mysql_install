package impl

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/solodba/mcube/logger"
)

// 解压MySQL压缩文件
func (i *impl) UnzipMySQLFile(ctx context.Context) error {
	_, err := os.Stat(i.c.MySQL.InstallPath)
	if err == nil {
		return fmt.Errorf("[%s]文件夹已经存在,请确定是否安装了MySQL", i.c.MySQL.InstallPath)
	}
	logger.L().Info().Msgf("正在解压MySQL压缩包")
	cmd := exec.Command("xd", "-d", i.c.MySQL.FileName)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("解压MySQL压程序失败: %v", err.Error())
	}
	tarName := strings.TrimRight(i.c.MySQL.FileName, ".xz")
	cmd = exec.Command("tar", "xf", tarName)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("解压MySQL压程序失败: %v", err.Error())
	}
	dirName := strings.TrimRight(tarName, ".tar")
	cmd = exec.Command("mv", dirName, i.c.MySQL.InstallPath)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("移动MySQL解压包失败: %v", err.Error())
	}
	logger.L().Info().Msgf("解压MySQL压缩包成功")
	return nil
}

// 创建MySQL相关目录
func (i *impl) CreateMySQLDir(ctx context.Context) error {
	return nil
}

// 判断是否有MySQL进程
func (i *impl) IsMySQLRun(ctx context.Context) error {
	return nil
}

// 创建MySQL用户
func (i *impl) CreateMySQLUser(ctx context.Context) error {
	return nil
}

// 修改权限
func (i *impl) ChangeMySQLDirPerm(ctx context.Context) error {
	return nil
}

// MySQL初始化
func (i *impl) InitialMySQL(ctx context.Context) error {
	return nil
}

// 启动MySQL
func (i *impl) StartMySQL(ctx context.Context) error {
	return nil
}
