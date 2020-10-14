package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Mode         string `mapstructure:"mode"`
	Port         string `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"db"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     string    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
}

var Conf = new(Config)

func Init() error {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {             // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 将读取的配置信息保存至全局变量Conf
	if err := viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}

	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被人修改啦...")
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
		}
	})

	return err
}
