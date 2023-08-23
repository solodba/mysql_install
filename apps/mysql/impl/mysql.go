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
	cmd := exec.Command("xz", "-d", i.c.MySQL.FileName)
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
	_, err := os.Stat(i.c.MySQL.DataBase)
	if err == nil {
		return fmt.Errorf("[%s]文件夹已经存在, 请确定是否安装了MySQL", i.c.MySQL.DataBase)
	}
	err = os.MkdirAll(i.c.MySQL.BinLogPath(), 0755)
	if err != nil {
		return fmt.Errorf("创建目录[%s]失败", i.c.MySQL.BinLogPath())
	}
	logger.L().Info().Msgf("创建目录[%s]成功", i.c.MySQL.BinLogPath())
	err = os.MkdirAll(i.c.MySQL.DataFilePath(), 0755)
	if err != nil {
		return fmt.Errorf("创建目录[%s]失败", i.c.MySQL.DataFilePath())
	}
	logger.L().Info().Msgf("创建目录[%s]成功", i.c.MySQL.DataFilePath())
	err = os.MkdirAll(i.c.MySQL.LogFilePath(), 0755)
	if err != nil {
		return fmt.Errorf("创建目录[%s]失败", i.c.MySQL.LogFilePath())
	}
	logger.L().Info().Msgf("创建目录[%s]成功", i.c.MySQL.LogFilePath())
	err = os.MkdirAll(i.c.MySQL.TmpFilePath(), 0755)
	if err != nil {
		return fmt.Errorf("创建目录[%s]失败", i.c.MySQL.TmpFilePath())
	}
	logger.L().Info().Msgf("创建目录[%s]成功", i.c.MySQL.TmpFilePath())
	err = os.MkdirAll(i.c.MySQL.ConfFilePath(), 0755)
	if err != nil {
		return fmt.Errorf("创建目录[%s]失败", i.c.MySQL.ConfFilePath())
	}
	logger.L().Info().Msgf("创建目录[%s]成功", i.c.MySQL.ConfFilePath())
	return nil
}

// 判断是否有MySQL进程
func (i *impl) IsMySQLRun(ctx context.Context) error {
	cmd := exec.Command("/bin/bash", "-c", `ps -ef |grep mysqld |wc -l`)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	if strings.Trim(string(res), "\n") != "2" {
		return fmt.Errorf("有MySQL进程在运行, 请检查")
	}
	return nil
}

// 创建MySQL用户
func (i *impl) CreateMySQLUser(ctx context.Context) error {
	cmd := exec.Command("/bin/bash", "-c", `cat /etc/passwd |grep -w mysql|wc -l`)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	if strings.Trim(string(res), "\n") != "1" {
		logger.L().Info().Msgf("MySQL用户不存在, 开始添加MySQL用户")
		cmd := exec.Command("groupadd", "mysql")
		_, err := cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
		cmd = exec.Command("useradd", "-g", "mysql", "mysql")
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
		logger.L().Info().Msgf("添加MySQL用户成功")
	}
	return nil
}

// 修改权限
func (i *impl) ChangeMySQLDirPerm(ctx context.Context) error {
	cmd := exec.Command("chown", "-R", "mysql:mysql", i.c.MySQL.DataBase)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	cmd = exec.Command("chown", "-R", "mysql:mysql", i.c.MySQL.InstallPath)
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	cmd = exec.Command("cp", "my.cnf", i.c.MySQL.ConfFilePath()+"/")
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	return nil
}

// MySQL初始化
func (i *impl) InitialMySQL(ctx context.Context) error {
	logger.L().Info().Msgf("开始初始化MySQL")
	cmdStr := fmt.Sprintf("%s/bin/mysqld --defaults-file=%s/my.cnf --user=mysql --initialize", i.c.MySQL.InstallPath, i.c.MySQL.ConfFilePath())
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	cmd = exec.Command("/bin/bash", "-c", `cat /data/mysql/log/mysql.err |grep -i "root@localhost:"|wc -l`)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	if strings.Trim(string(res), "\n") != "1" {
		return fmt.Errorf("MySQL初始化失败")
	}
	logger.L().Info().Msgf("MySQL初始化成功")
	return nil
}

// 启动MySQL
func (i *impl) StartMySQL(ctx context.Context) error {
	return nil
}
