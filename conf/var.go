package conf

import (
	"embed"
	"time"
)

//go:embed *.yaml
var ConfYamlDir embed.FS

type AppConf struct {
	AppName                             string           `yaml:"app-name"`
	IsUseRabbitMq                       bool             `yaml:"isUseRabbitMq"`
	CommonGoRoutinePoolSize             int              `yaml:"commonGoRoutinePoolSize"`
	CommonGoRoutinePoolMinuteExpire     int              `yaml:"commonGoRoutinePoolMinuteExpire"`
	CommonGoRoutinePoolMaxBlockingTasks int              `yaml:"commonGoRoutinePoolMaxBlockingTasks"`
	AccessTokenTimeout                  int64            `yaml:"accessTokenTimeout"`
	RefreshTokenTimeout                 int64            `yaml:"refreshTokenTimeout"`
	CaptchaTimout                       int64            `yaml:"captchaTimout"`
	SmsCodeTimout                       int64            `yaml:"smsCodeTimout"`
	IsDemoSite                          bool             `yaml:"isDemoSite"`
	Ssl                                 bool             `yaml:"ssl"`
	Refer                               []string         `yaml:"refer"`
	Env                                 string           `yaml:"env"`
	Server                              ServerCfg        `yaml:"server"`
	Redis                               RedisCfg         `yaml:"redis"`
	Mysql                               MysqlCfg         `yaml:"mysql"`
	Log                                 LogCfg           `yaml:"log"`
	RabbitMQ                            RabbitMQCfg      `yaml:"rabbitmq"`
	Elasticsearch                       ElasticsearchCfg `yaml:"elasticsearch"`
}

type ServerCfg struct {
	RunMode      string        `yaml:"runMode"`
	HttpPort     int           `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type MysqlCfg struct {
	Base struct {
		MaxOpenConn     int           `yaml:"maxOpenConn"`
		MaxIdleConn     int           `yaml:"maxIdleConn"`
		ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
	} `yaml:"base"`

	Read struct {
		Addr string `yaml:"addr"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"read"`
	Write struct {
		Addr string `yaml:"addr"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"write"`
}

type RedisCfg struct {
	Addr        string `yaml:"addr"`
	Db          int    `yaml:"db"`
	MaxRetries  int    `yaml:"maxRetries"`
	MinIdleConn int    `yaml:"minIdleConn"`
	Pass        string `yaml:"pass"`
	PoolSize    int    `yaml:"poolSize"`
}

type LogCfg struct {
	AppName     string `yaml:"app-name"`
	LogSavePath string `yaml:"logSavePath"`
	TimeFormat  string `yaml:"timeFormat"`
}

type RabbitMQCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type ElasticsearchCfg struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}
