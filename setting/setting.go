package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AppConfig struct {
	Port int    `yaml:"port"`
	Mode string `yaml:"mode"`
	Ver  string `yaml:"ver"`
}

type MySQLConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DB           string `yaml:"db"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	MaxConns     int    `yaml:"max_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

type RedisConfig struct {
	DB   string `yaml:"db"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func Init() error {
	// 从配置文件中加载配置信息
	viper.SetConfigName("config")  // 指定配置文件名称（不需要带后缀）
	viper.SetConfigType("yaml")    // 指定配置文件类型
	viper.AddConfigPath(".")       // 指定查找配置文件的路径（这里使用相对路径）
	viper.AddConfigPath("./conf/") // 指定查找配置文件的路径（这里使用相对路径）
	err := viper.ReadInConfig()    // 读取配置信息
	viper.Unmarshal(&Conf)
	viper.WatchConfig() // 监控配置文件的变更
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		viper.Unmarshal(&Conf) // 当配置文件修改之后要把变更后的配置信息更新到全局变量Conf里
		fmt.Println("xxxx, 配置文件被人修改啦！:", e.Name)
	})
	return err
}
