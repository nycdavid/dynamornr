package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func main() {
	dbyml := make(map[interface{}]interface{})
	fpath, err := filepath.Abs("./config/database.yml")
	if err != nil {
		log.Fatal(err)
	}
	file, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(file, &dbyml)
	envConfigs := dbyml["test"].(map[interface{}]interface{})
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "Fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		fmt.Printf("Test database is running on %s:%d\n", envConfigs["host"], envConfigs["port"])
		return nil
	}
	app.Run(os.Args)
}
