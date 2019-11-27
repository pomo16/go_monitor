package model

type Config struct {
	Mysql DBConfig `yaml:"mysql"`
}

type DBConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
