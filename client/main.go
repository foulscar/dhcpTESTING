package main

import (
	"log"
	"os"
)

func main() {
	cfgFile := "defaultConfig.ini"
	if len(os.Args) > 1 {
		cfgFile = os.Args[1]
	}

	cfg, err := ParseConfigFile(cfgFile)
	if err != nil {
		log.Fatalln("Error parsing config. ", err.Error())
	}

	rt, err := cfg.BuildRuntime()
	if err != nil {
		log.Fatalln("Error building runtime. ", err.Error())
	}
        defer rt.Conn.Close()
        
	rt.PromptMainMenu()
}
