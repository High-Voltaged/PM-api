package config

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
	Database struct {
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Port     string `mapstructure:"port"`
		Host     string `mapstructure:"host"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"db"`
}
