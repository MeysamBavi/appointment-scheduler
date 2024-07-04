package config

import (
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/httpserver"
	"github.com/MeysamBavi/appointment-scheduler/backend/pkg/postgres"
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
	CORS           httpserver.CORSConfig `config:"cors"`
	Postgres       postgres.Config       `config:"postgres"`
	TheWallAddress string                `config:"the_wall_address"`
}

func Default() Config {
	return Config{
		CORS: httpserver.CORSConfig{
			Enable:  true,
			Origins: nil,
		},
		Postgres: postgres.Config{
			Host:               "postgres",
			Port:               "5432",
			User:               "postgres",
			Password:           "postgres",
			DBName:             "appointment_scheduler",
			MaxIdleConnections: 1,
			MaxConnections:     5,
		},
		TheWallAddress: "http://the-wall",
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
