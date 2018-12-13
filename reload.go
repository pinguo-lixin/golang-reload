package reload

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
)

// Reloader reload interface
type Reloader interface {
	Reload() error
}

// Config reload config
type Config struct {
	PidFile string
	Signal  syscall.Signal
}

// DefaultSignal default signal trigger reload
var DefaultSignal = syscall.SIGUSR2

// NewConfig init config
func NewConfig(pidFile string, signal syscall.Signal) Config {
	return Config{PidFile: pidFile, Signal: signal}
}

// Listen signal listen
func (c Config) Listen(reloader Reloader) error {
	if err := c.savePid(); err != nil {
		return err
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, c.Signal)

	go func() {
		for {
			<-s
			if err := reloader.Reload(); err != nil {
				panic(err)
			}
		}
	}()

	return nil
}

// Reload reload action
func (c Config) Reload() error {
	pid, err := c.getPid()
	if err != nil {
		return err
	}

	s := strconv.Itoa(int(c.Signal))

	return exec.Command("kill", "-"+s, pid).Run()
}

func (c Config) savePid() error {
	pid := os.Getpid()
	return ioutil.WriteFile(c.PidFile, []byte(strconv.Itoa(pid)), 0666)
}

func (c Config) getPid() (string, error) {
	buf, err := ioutil.ReadFile(c.PidFile)
	return string(buf), err
}
