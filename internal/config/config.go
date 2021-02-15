package conf

import (
	sq "github.com/goclub/sql"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type ConfigKVDS struct {
	Network string `yaml:"network"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
type Config struct {
	RDS sq.DataSource `yaml:"rds"`
	KVDS ConfigKVDS `yaml:"kvds"`
}
func NewConfig () (config Config, err error) {
	gopath := os.Getenv("GOPATH")// 获取系统的环境变量
	projectPath := path.Join(gopath, "src/github.com/goclub/project-seckilling")
	envPath := path.Join(projectPath, "env/env.yaml")
	data, err := ioutil.ReadFile(envPath) ; if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &config) ; if err != nil {
		return
	}
	return
}

func (c *Config) Check() error {
	return nil
}



func TestConfig (t *testing.T) (testConfig Config) {
	var err error
	testConfig, err = NewConfig();if err != nil {
		panic(err)
	}
	return
}