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
	cmd := exec.Command("/bin/bash", "-c", `ps -ef | grep mysqld | grep -v grep | wc -l`)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	if strings.Trim(string(res), "\n") != "0" {
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
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	cmd = exec.Command("cp", "my.cnf", i.c.MySQL.ConfFilePath()+"/")
	_, err = cmd.Output()
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
	cmdStr := fmt.Sprintf(`cat %s/mysql.err | grep 'temporary password'`, i.c.MySQL.LogFilePath())
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	pwdList := strings.Split(string(res), " ")
	pwd := strings.TrimRight(pwdList[len(pwdList)-1], "\n")
	_, err = os.Stat("/etc/init.d/mysql.server")
	if err != nil {
		cmd := exec.Command("cp", "mysql.server", "/etc/init.d/", "-rf")
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
		cmd = exec.Command("chmod", "700", "/etc/init.d/mysql.server")
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
	}
	err = i.AddEnv(ctx)
	if err != nil {
		return err
	}
	cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "/etc/init.d/mysql.server start > /dev/null 2>&1")
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	cmdStr = fmt.Sprintf(`source /etc/profile;mysql -uroot -p'%s' --connect-expired-password -e "alter user user() identified by 'Root@123';"`, pwd)
	cmd = exec.Command("/bin/bash", "-c", cmdStr)
	_, err = cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	logger.L().Info().Msgf("MySQL 8.0.25 安装完成")
	return nil
}

// 增加环境量变量
func (i *impl) AddEnv(ctx context.Context) error {
	cmdStr := fmt.Sprintf(`grep 'export PATH=$PATH:%s/bin' /etc/profile|wc -l`, i.c.MySQL.InstallPath)
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	res, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("执行命令失败, err: %s", err.Error())
	}
	if strings.Trim(string(res), "\n") == "0" {
		cmdStr = fmt.Sprintf(`echo "export PATH=\$PATH:%s/bin" >> /etc/profile`, i.c.MySQL.InstallPath)
		cmd := exec.Command("/bin/bash", "-c", cmdStr)
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
		cmd = exec.Command("/bin/bash", "-c", `source /etc/profile`)
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("执行命令失败, err: %s", err.Error())
		}
	}
	return nil
}
