package databases

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type databaseYml struct {
	env   string
	fpath string
}

func newDatabaseYml(ctx *cli.Context) *DatabaseYml {
	configPath := parseConfigPath(ctx)
	fpath := filepath.Join(configPath, "database.yml")
	return &databaseYml{
		env:         os.Getenv("ENV"),
		fileContent: fileContents(fpath),
	}
}

func (dbyml *databaseYml) url() string {
	dbYml := make(map[interface{}]interface{})
	yaml.Unmarshal(dbyml.fileContent, &dbYml)
	envConfigs := dbYml[dbyml.env].(map[interface{}]interface{})
	return fmt.Sprintf("%s:%v", envConfigs["host"], envConfigs["port"])
}

func fileContents(path) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	return content
}

func parseConfigPath(ctx *cli.Context) string {
	if ctx.String("config") == "" {
		return "config"
	}
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	path := fmt.Sprintf("%s/%s", projectDir, ctx.String("config"))
	return path
}
