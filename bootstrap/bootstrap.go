package bootstrap

import (
	"github.com/davveo/lemonShop-framework/cache"
	"github.com/davveo/lemonShop-framework/db"
	"github.com/davveo/lemonShop-framework/logger"
	"github.com/davveo/lemonShop-seller-server/app"
	"github.com/davveo/lemonShop-seller-server/conf"
)

func Bootstrap() {
	appConf, err := conf.Init()
	if err != nil {
		panic(err)
	}
	gLogger, err := logger.Init(&logger.LogCfg{
		AppName:     appConf.Log.AppName,
		LogSavePath: appConf.Log.LogSavePath,
		TimeFormat:  appConf.Log.TimeFormat,
	})
	if err != nil {
		panic(err)
	}

	cacheRepo, err := cache.Init(&cache.RedisCfg{
		Addr:        appConf.Redis.Addr,
		Db:          appConf.Redis.Db,
		MaxRetries:  appConf.Redis.MaxRetries,
		MinIdleConn: appConf.Redis.MinIdleConn,
		Pass:        appConf.Redis.Pass,
		PoolSize:    appConf.Redis.PoolSize,
	})
	if err != nil {
		panic(err)
	}

	dbRepo, err := db.Init(&db.MysqlCfg{
		Base: db.MysqlBase{
			OpenSlaveRead:   false,
			MaxOpenConn:     appConf.Mysql.Base.MaxOpenConn,
			MaxIdleConn:     appConf.Mysql.Base.MaxIdleConn,
			ConnMaxLifeTime: appConf.Mysql.Base.ConnMaxLifeTime,
		},
		Write: db.MysqlIns{
			Addr: appConf.Mysql.Write.Addr,
			User: appConf.Mysql.Write.User,
			Pass: appConf.Mysql.Write.Pass,
			Name: appConf.Mysql.Write.Name,
		},
		Read: db.MysqlIns{
			Addr: appConf.Mysql.Write.Addr,
			User: appConf.Mysql.Write.User,
			Pass: appConf.Mysql.Write.Pass,
			Name: appConf.Mysql.Write.Name,
		},
	})
	if err != nil {
		panic(err)
	}

	server := app.NewServer(appConf, gLogger, dbRepo, cacheRepo)
	server.Init()
}
