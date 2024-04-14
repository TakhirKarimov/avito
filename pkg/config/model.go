package config

type Config struct {
	DBConfig    DBConfig    `mapstructure:"database"`
	CacheConfig CacheConfig `mapstructure:"cache_database"`
}

type DBConfig struct {
	Dsn string `mapstructure:"dsn"`
}

type CacheConfig struct {
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
