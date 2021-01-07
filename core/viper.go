package core

import (
	"flag"
	"fmt"
	"go-svc-bootstrap/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // priority: CLI > env > default value
			if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" {
				config = global.ConfigFile
				fmt.Printf("You are using the default config path value, and the config path is %v\n", global.ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("You are using the GVA_CONFIG env, and the config path is %v\n", config)
			}
		} else {
			fmt.Printf("You are using the CLI -c param value, and the config path is %v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("You are using the func Viper() param value, and the config path is %v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.G_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&global.G_CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
