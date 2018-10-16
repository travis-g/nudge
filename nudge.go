package nudge

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

var (
	PID int

	flagDisabled bool
	flagInterval time.Duration
	flagKey      string
)

// Nudge keeps the computer awake by pressing a provided key. It's simply a
// wrapper for `robotgo.KeyTap()`.
func Nudge(key string, modifiers interface{}) {
	robotgo.KeyTap(key, modifiers)
}

// Logs to a file (stdout, stderr, or file)
func log(s string, file *os.File) {
	fmt.Fprintf(file, "%s\n", s)
}

// Parse any flags sent to the command
func parseFlags() {
	flag.BoolVar(&flagDisabled, "disable", false, "start disabled")
	flag.DurationVar(&flagInterval, "interval", 59*time.Second, "interval between nudges")
	flag.StringVar(&flagKey, "key", "f16", "key to press")
	flag.Parse()
}

func main() {
	parseFlags()
	PID = os.Getpid()
	log(fmt.Sprintf("pid: %d", PID), os.Stderr)

	// Set interrupt keys:
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	// Setup toggling of disabled status via SIGUSR1:
	sigToggle := make(chan os.Signal, 0)
	signal.Notify(sigToggle, syscall.SIGUSR1)
	go func() {
		// Set up an infinite loop to continually listen for SIGUSR1 and toggle
		// once per signal:
		for {
			<-sigToggle
			flagDisabled = !flagDisabled
			log(fmt.Sprintf("disabled: %t", flagDisabled), os.Stdout)
		}
	}()

	// Set up the nudger:
	nudger := time.NewTicker(flagInterval)
	defer nudger.Stop()
	go func() {
		// Nudge at every interval:
		for t := range nudger.C {
			if !flagDisabled {
				Nudge(flagKey, nil)
				log(fmt.Sprintf("nudged: %v", t.Local().Format("15:04:05")), os.Stderr)
			}
		}
	}()

	log("send SIGINT/SIGTERM to quit, SIGUSR1 to toggle...", os.Stderr)

	// Block/keep running goroutine until a quit signal is received:
	s := <-sigQuit
	log(fmt.Sprintf("%v received", s), os.Stderr)
}
