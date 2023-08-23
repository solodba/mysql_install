package conf

import (
	"fmt"

	"github.com/solodba/mcube/logger"
)

// 全局配置参数
var (
	c *Config
)

func C() *Config {
	if c == nil {
		logger.L().Panic().Msgf("please initial global config")
	}
	return c
}

// 全局配置结构体
type Config struct {
	MySQL *MySQL `toml:"mysql"`
}

// MySQL安装配置结构体
type MySQL struct {
	FileName       string `toml:"file_name" env:"MYSQL_FILE_NAME"`
	InstallPath    string `toml:"install_path" env:"MYSQL_INSTALL_PATH"`
	DataBase       string `toml:"data_base" env:"MYSQL_DATA_BASE"`
	BinlogFileName string `toml:"binlog_file_name" env:"MYSQL_BINLOG_FILE_NAME"`
	DataFileName   string `toml:"data_file_name" env:"MYSQL_DATA_FILE_NAME"`
	LogFileName    string `toml:"log_file_name" env:"MYSQL_LOG_FILE_NAME"`
	TmpFileName    string `toml:"tmp_file_name" env:"MYSQL_TMP_FILE_NAME"`
	ConfFileName   string `toml:"conf_file_name" env:"MYSQL_CONF_FILE_NAME"`
}

// MySQL结构体初始化函数
func NewDefaultMySQL() *MySQL {
	return &MySQL{}
}

// Config结构体初始化函数
func NewDefaultConfig() *Config {
	return &Config{
		MySQL: NewDefaultMySQL(),
	}
}

// MySQL结构体添加方法
func (m *MySQL) BinLogPath() string {
	return fmt.Sprintf("%s/%s", m.DataBase, m.BinlogFileName)
}

func (m *MySQL) DataFilePath() string {
	return fmt.Sprintf("%s/%s", m.DataBase, m.DataFileName)
}

func (m *MySQL) LogFilePath() string {
	return fmt.Sprintf("%s/%s", m.DataBase, m.LogFileName)
}

func (m *MySQL) TmpFilePath() string {
	return fmt.Sprintf("%s/%s", m.DataBase, m.TmpFileName)
}

func (m *MySQL) ConfFilePath() string {
	return fmt.Sprintf("%s/%s", m.DataBase, m.ConfFileName)
}
