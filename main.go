package main

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/drivernation/kaiju"
	"os"
)

const LoggingConfig = `
<seelog>
    <outputs formatid="main">
        <console/>
    </outputs>
    <formats>
        <format id="main" format="%Date %Time [%Level]: %Msg%n"/>
    </formats>
</seelog>
`

var logger seelog.LoggerInterface
var configFile string

func init() {
	flag.StringVar(&configFile, "config", "", "The configuration file to use.")
	var err error
	logger, err = seelog.LoggerFromConfigAsString(LoggingConfig)
	if err != nil {
		panic(err)
	}

	kaiju.UseLogger(logger)
}

func main() {
	flag.Parse()
	if configFile == "" {
		logger.Error("No configuration file provided.")
		// Unfortunately, deferred functions are not executed when os.Exit(int) is called. T_T
		logger.Close()
		os.Exit(1)
	}

	var err error
	// LoadConfig(string) is assumed to be defined elsewhere.
	config, err := LoadConfigYaml(configFile)
	if err != nil {
		logger.Errorf("Failed to load configuration: %s", err)
		logger.Close()
		os.Exit(1)
	}

	helloHandler := HelloHandler{Saying: config.Saying}
	kaiju.Handle("/hello", &helloHandler, "GET")
	defer logger.Close()
	if err = kaiju.Start(config.Config); err != nil {
		panic(logger.Error(err))
	}
}
