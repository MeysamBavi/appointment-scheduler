package config

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"log"
)

const (
	tag            = "config"
	delimiter      = "."
	configFilePath = "/etc/config/config.yaml"
)

type Config struct {
	CORS httpserver.CORSConfig `config:"cors"`
}

func Default() Config {
	return Config{
		CORS: httpserver.CORSConfig{
			Enable:  true,
			Origins: nil,
		},
	}
}

var k = koanf.New(delimiter)

func Load() Config {
	{
		err := k.Load(structs.Provider(Default(), tag), nil)
		if err != nil {
			log.Fatalf("could not load default config: %v\n", err)
		}
	}

	{
		err := k.Load(file.Provider(configFilePath), yaml.Parser())
		if err != nil {
			log.Printf("could not load yaml config: %v\n", err)
		}
	}

	var instance Config
	err := k.UnmarshalWithConf("", &instance, koanf.UnmarshalConf{
		Tag: tag,
	})

	if err != nil {
		log.Fatalf("could not unmarshal config: %v\n", err)
	}

	return instance
}
