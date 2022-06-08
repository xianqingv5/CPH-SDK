package conf

import (
	"fmt"
	"os"

	"github.com/xuanxinhuiqing/gotools"
)

// 全局配置变量
var Config Conf

// 全局配置文件
type Conf struct {
	Huawei Huawei `json:"huawei"`
}

// url 配置
type Huawei struct {
	Endpoint  string `json:"endpoint"`
	ProjectId string `json:"project_id"`
	Key       string `json:"key"`
	Secret    string `json:"secret"`
}

// 加载配置文件
func Init() {
	env := os.Getenv("IS_ENV")
	if env == "dev" { // 开发环境
		fmt.Println("[env]", "开发环境 --", env)
		if err := gotools.DecodeJsonFile("conf/cph.json", &Config); err != nil {
			fmt.Println(err)
			panic(err)
		}
	} else { // 生产环境
		fmt.Println("[env]", "生产环境 --", env)
		if err := gotools.DecodeJsonFile("conf/cph.json", &Config); err != nil {
			panic(err)
		}
	}

}
