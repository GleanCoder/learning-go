package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string
}

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-required:"true"`
	Storage    string `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server"`
}

// this function should execute successfully, if any error occured in this func, then we will not start our application further.
func MustLoad() *Config {
	// step-1: Get config path.
	var configPath string

	// when we deploye the config can be anywhere in our deployed machine, so we have to pass the path of that.
	// we can pass it in 2 way: one we can pass it through env variables, to get that we have to use os getenv method

	configPath = os.Getenv("CONFIG_PATH")

	//step-2: check if path is present or not

	// then we will check if it's empty then anyone pass it through termanal arguments during running the program which we pass in commands
	// go run main.go -config-path xyz even if the flags were not present then we will throw fatal error to not go further in application
	if configPath == "" {
		flags := flag.String("config", "", "path to the config")
		flag.Parse()
		// then pass this to configPath

		configPath = *flags // as we see it's a pointer so we have to dereference it for the value

		// if path is empty then throw fatal error and stop application

		if configPath == "" {
			log.Fatal("Config path is missing")
		}

	}
	//step-3: check if there any file present on that path if not then throw fatal error to stop the application

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file is not present %s", configPath)
	}

	//step-4: if all the above path and file is present then our application will go further

	var cfg Config

	// to parse config we can use cleanenv here

	err := cleanenv.ReadConfig(configPath, &cfg) // so here we have to pass the path from where it'll read the file and have to pass the struct where it will parse the struct tags

	if err != nil {
		log.Fatalf("cann't read config file %s", err.Error())
	}

	return &cfg

}
