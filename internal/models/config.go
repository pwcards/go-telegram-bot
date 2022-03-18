package models

type Config struct {
	Telegram struct {
		Token string `yaml:"token"`
	} `yaml:"telegram"`
	DataBase   DBConfig   `yaml:"mysql"`
	ServerData ServerData `yaml:"server_data"`
}

type DBConfig struct {
	Host               string `mapstructure:"HOST" valid:"required"`
	Port               string `mapstructure:"PORT" valid:"required"`
	Username           string `mapstructure:"USERNAME" valid:"required"`
	Password           string `mapstructure:"PASSWORD" valid:"required"`
	DBName             string `mapstructure:"DBNAME" valid:"required"`
	OnlyFastQueries    bool   `mapstructure:"ONLY_FAST_QUERIES" valid:"-"`
	ConnectRetryTTL    int    `mapstructure:"CONNECT_RETRY_TTL" valid:"-"`
	ConnectRetryMAX    int    `mapstructure:"CONNECT_RETRY_MAX" valid:"-"`
	ConnectCheckTTL    int    `mapstructure:"CONNECT_CHECK_TTL" valid:"required"`
	MaxOpenConnections int    `mapstructure:"MAX_OPEN_CONN" valid:"-"`
	MaxIdleConnections int    `mapstructure:"MAX_IDLE_CONN" valid:"-"`
}

type ServerData struct {
	Host string `mapstructure:"HOST" valid:"required"`
}
