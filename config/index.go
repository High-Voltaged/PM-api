package config

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"db"`
	Email    Email    `mapstructure:"email"`
}

type Server struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
}

type Database struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Name     string `mapstructure:"name"`
}

type Email struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	From     string `mapstructure:"from"`
}
