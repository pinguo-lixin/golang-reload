package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/pinguo-lixin/golang-reload"
	"syscall"
	"time"
)

// Config config items
type Config struct {
	Mode string
}

func loadConfig() error {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(f, &config)
}

type reloaderHandler struct{}

func (r *reloaderHandler) Reload() error {
	return loadConfig()
}

var (
	config *Config
	cmd    string // start || reload

	reloader = reload.NewConfig("pid", syscall.SIGUSR2)
)

func init() {
	if len(os.Args) < 2 {
		cmd = "start"
	} else {
		cmd = os.Args[1]
	}
}

func main() {
	if cmd == "start" {
		start()
	} else if cmd == "reload" {
		reloader.Reload()
	} else {
		log.Panicf("unknow command %s\n", cmd)
	}
}

func start() {

	if err := loadConfig(); err != nil {
		panic(err)
	}

	r := new(reloaderHandler)
	if err := reloader.Listen(r); err != nil {
		panic(err)
	}

	for {
		time.Sleep(2 * time.Second)
		fmt.Printf("mode: %s\n", config.Mode)
	}
}
