package config

// database struct
type Database struct {
	Username    string `yaml:"postgres.username" required:"true"`
	Password    string `yaml:"postgres.password" required:"true"`
	Host        string `yaml:"postgres.host" required:"true"`
	Port        string `yaml:"postgres.port" required:"true"`
	Schema      string `yaml:"postgres.schema" required:"true"`
	Automigrate bool   `yaml:"postgres.automigrate"`
	Logger      bool   `yaml:"postgres.logger"`
	Namespace   string `yaml:"postgres.namespace"`
}
