package config

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       string `yaml:"DB"`
}

func (r RedisConfig) GetAddr() string {
	return r.Addr
}

func (r RedisConfig) GetPwd() string {
	return r.Password
}
