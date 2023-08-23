package conf_test

import (
	"testing"

	"github.com/solodba/mysql_install/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	err := conf.LoadConfigFromToml("test/test.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().MySQL)
	t.Log(conf.C().MySQL.BinLogPath())
	t.Log(conf.C().MySQL.DataFilePath())
	t.Log(conf.C().MySQL.LogFilePath())
	t.Log(conf.C().MySQL.TmpFilePath())
	t.Log(conf.C().MySQL.ConfFilePath())
}

func TestLoadConfigFromEnv(t *testing.T) {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(conf.C().MySQL)
	t.Log(conf.C().MySQL.BinLogPath())
	t.Log(conf.C().MySQL.DataFilePath())
	t.Log(conf.C().MySQL.LogFilePath())
	t.Log(conf.C().MySQL.TmpFilePath())
	t.Log(conf.C().MySQL.ConfFilePath())
}
