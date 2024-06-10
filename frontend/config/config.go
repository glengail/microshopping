package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		DatabaseSettings
		JWTSettings
	}
	DatabaseSettings struct {
		DatabaseURI  string
		DatabaseName string
		Username     string
		Password     string
	}
	JWTSettings struct {
		SecretKey string
	}
	ConfigReader struct {
		configFile string
		v          *viper.Viper
	}
)

func GetAllConfigValues(configFile string) (configuration *Configuration, err error) {
	configReader := newConfigReader(configFile)
	err = configReader.v.ReadInConfig()
	if err != nil {
		fmt.Printf("读取配置文件失败：%s\n", err)
		return nil, err
	}
	//解析配置文件到结构体
	err = configReader.v.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("解析配置文件失败：%s\n", err)
		return nil, err
	}
	return configuration, err
}
func newConfigReader(configFile string) *ConfigReader {
	v := viper.GetViper()
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")
	return &ConfigReader{
		configFile: configFile,
		v:          v,
	}
}
