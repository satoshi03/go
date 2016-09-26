package config

type Config struct {
	Redis  map[string]Redis `yaml:"redis"`
	Fluent Fluent           `yaml:"fluent"`
	Cookie Cookie           `yaml:"cookie"`
}

type Redis struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DB     int    `yaml:"db"`
	Slaves []RedisConfig
	Option RedisOption
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	DB   int    `yaml:"db"`
}

type RedisOption struct {
	TTL int `yaml:"ttl"`
}

type Fluent struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Cookie struct {
	Name   string `yaml:"name"`
	Domain string `yaml:"domain"`
}
