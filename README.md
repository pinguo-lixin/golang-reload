# Golang reloader

A reloader component for golang project.

Useage:
```go

func loadConfig() error {
    // 
}

type reloadHandler struct{}

func (r *reloadHandler) Reload() error {
	return loadConfig()
}

var reloader = reload.NewConfig("pid", syscall.SIGUSR2)

func main() {
    var cmd = "start"

    if len(os.Args) > 1 && os.Args[1] == "reload" {
        cmd = "reload"
    }

    if cmd == "start" {
        start()
    } else if cmd == "reload" {
        reloader.Reload()
    }
}

func start() {
    r := new(reloadHandler)
	if err := reloader.Listen(r); err != nil {
		panic(err)
    }
    
    // 
}
```
