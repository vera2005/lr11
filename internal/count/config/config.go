package config

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`

	API     api     `yaml:"api"`
	Usecase usecase `yaml:"usecase"`
	DB      db      `yaml:"db"`
}

type api struct {
	MaxNum int `yaml:"max_number"`
}

type usecase struct {
	DefaultCount string `yaml:"default_count"`
}

type db struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBname   string `yaml:"dbname"`
}
