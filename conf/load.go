package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
)

// 通过Toml文件加载全局配置
func LoadConfigFromToml(filePath string) error {
	c = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, c)
	if err != nil {
		return err
	}
	return nil
}

// 通过环境变量文件加载全局配置
func LoadConfigFromEnv() error {
	c = NewDefaultConfig()
	return env.Parse(c)
}
