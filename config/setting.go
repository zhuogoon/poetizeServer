package config

type SettingConfig struct {
	MysqlConfig MysqlConfig `yaml:"mysql"`
	RedisConfig RedisConfig `yaml:"redis"`
}
