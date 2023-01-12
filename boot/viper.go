package boot

import (
	"CSAwork/global"
	"fmt"
	"github.com/spf13/viper"
)

func ViperSetup(configPath string) {
	v := viper.New()
	v.SetConfigFile(configPath) // 设置配置文件路径
	v.SetConfigType("yaml")     // 设置配置文件类型
	err := v.ReadInConfig()     // 读取配置文件
	if err != nil {
		panic(fmt.Errorf("get config file failed, err: %v", err))
	}

	if err = v.Unmarshal(&global.Config); err != nil {
		// 将配置文件反序列化到 Config 结构体
		panic(fmt.Errorf("unmarshal config failed, err: %v", err))
	}
}
