package common

import "github.com/BurntSushi/toml"

type Configs struct {
	Listen int
	MySQL  MySQLConfig
}

type MySQLConfig struct {
	DSN     string
	MaxIdle int
	MaxOpen int
}

// Config 全局配置信息
var Config *Configs

// InitConfig 加载配置
func InitConfig(path string) {
	config, err := loadConfig(path)
	if err != nil {
		panic(err)
	}
	Config = config
}

func loadConfig(path string) (*Configs, error) {
	config := new(Configs)
	if _, err := toml.DecodeFile(path, config); err != nil {
		return nil, err
	}

	return config, nil
}
