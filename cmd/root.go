package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/op/go-logging"
)

var (
	cfgFile string
	logFilename string
	log = logging.MustGetLogger("default")
)

type LogWriter struct {
	logFunc	func(format string, args ...interface{})
}

func NewLogWriter(level logging.Level) (lw LogWriter) {
	lw = LogWriter{}
	switch level {
	case logging.CRITICAL:
		lw.logFunc = log.Criticalf
	case logging.ERROR:
		lw.logFunc = log.Errorf
	case logging.WARNING:
		lw.logFunc = log.Warningf
	case logging.NOTICE:
		lw.logFunc = log.Noticef
	case logging.INFO:
		lw.logFunc = log.Infof
	case logging.DEBUG:
		lw.logFunc = log.Debugf
	}
	return lw
}

func (lw LogWriter) Write(data []byte) (n int, err error) {
	lw.logFunc("%s", data)
	return len(data), nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sprk",
	Short: "Sparky is a good dog and can do all kinds of tricks",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		fileBackend := logging.NewLogBackend(logFile, "", 0)
		stdoutBackend := logging.AddModuleLevel(logging.NewLogBackend(os.Stdout, "", 0))
		stdoutBackend.SetLevel(logging.INFO, "")
		logging.SetFormatter(logging.MustStringFormatter("%{color}%{time:15:04:05.000} ▶ %{level:.10s} %{color:reset} %{message}"))
		logging.SetBackend(fileBackend, stdoutBackend)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logFilename, "log", "sparky.log", "File where logs will be written")
}
