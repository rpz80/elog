# elog
Really simple logging library for Go. Prints your messages to stderr.

# Usage
1. Pass `-log-level=LEVEL` to your executable. `DEBUG`, `INFO`, `WARNING`, `ERROR`, `CRITICAL` are
supported.
2. Call `flag.Parse()` in `main()`.

# Example
```
import (
	"flag"

	"github.com/rpz80/elog"
)

func main() {
	flag.Parse()
	intVar := 42
	stringVar := "Really valuable string"
	elog.Debug("Something not that important, %v", intVar)
	elog.Info("Service started: %v", stringVar)
	elog.Warning("Something not that expected")
	elog.Error("This is clearly an error")
	elog.Critical("Crashing") // panic("Crashing") will be called here
}
```
