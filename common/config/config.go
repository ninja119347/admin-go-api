package config

import (
	//"io/ioutil"
	"gopkg.in/yaml.v2"
	"os"
)

type config struct {
	Server        server        `yaml:"server"`
	DB            db            `yaml:"db"`
	Redis         redis         `yaml:"redis"`
	Log           log           `yaml:"log"`
	ImageSettings imageSettings `yaml:"imageSettings"`
}
type server struct {
	Address string `yaml:"address"`
	Model   string `yaml:"model"`
}

// 数据库配置
type db struct {
	Dialects string `yaml:"dialects"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"db"`
	//DB       string `yaml:"db"`
	Charset string `yaml:"charset"`
	MaxIdle int    `yaml:"maxIdle"`
	MaxOpen int    `yaml:"maxOpen"`
}

// redis settings
type redis struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

// picture settings
type imageSettings struct {
	UploadDir string `yaml:"uploadDir"`
	ImageHost string `yaml:"imageHost"`
}

// log settings
type log struct {
	Path  string `yaml:"path"`
	Name  string `yaml:"name"`
	Model string `yaml:"model"`
}

var Config *config

func init() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err1 := yaml.Unmarshal(file, &Config)
	if err1 != nil {
		panic(err1)
	}
}
