package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gvb/config"
	"gvb/global"
	"log"
	"os"
)

// InitConfig 读取yaml配置
func InitConfig() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("获取yaml文件失败: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success")
	global.Config = c
}
