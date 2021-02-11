package conf

import (
	sq "github.com/goclub/sql"
	"gopkg.in/yaml.v2"
	_ "gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
)

type Config struct {
	RDS sq.DataSource `yaml:"rds"`
}
func NewConfig () (config Config, err error) {
	gopath := os.Getenv("GOPATH")
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