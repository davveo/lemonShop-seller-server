package conf

import (
	"github.com/davveo/lemonShop-framework/config"
	"os"
)

var Conf AppConf

const (
	activeProfiles = "active"
)

func Init() (*AppConf, error) {
	err := config.Unmarshal(ConfYamlDir,
		&Conf, config.ProfileFunc(func() string {
			active := os.Getenv(activeProfiles)
			if active == "" {
				active = "dev" // 默认配置
			}
			return active
		}))
	if err != nil {
		return nil, err
	}
	return &Conf, nil
}
