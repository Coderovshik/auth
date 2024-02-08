package config

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		TokenTTL time.Duration `env:"TOKEN_TTL" env-default:"1h"`
	}
	GRPC struct {
		Port string `env:"GRPC_PORT" env-default:"50051"`
		Host string `env:"GRPC_HOST" env-default:"localhost"`
	}
	MongoDB struct {
		Port     string `env:"MONGO_PORT" env-required:"true"`
		Host     string `env:"MONGO_HOST" env-required:"true"`
		Username string `env:"MONGO_USERNAME" env-required:"true"`
		Password string `env:"MONGO_PASSWORD" env-required:"true"`
		Database string `env:"MONGO_DATABASE_NAME" env-required:"true"`
	}
	Migrations struct {
		MigrationsCollection string `env:"MIGRATIONS_COLLECTION" env-default:"migrations"`
		MigrationsPath       string `env:"MIGRATIONS_PATH" env-required:"true"`
	}
}

func (c *Config) AddressGRPC() string {
	return net.JoinHostPort(c.GRPC.Host, c.GRPC.Port)
}

func (c *Config) MongoURI() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s",
		c.MongoDB.Username,
		c.MongoDB.Password,
		c.MongoDB.Host,
		c.MongoDB.Port,
	)
}

func _get() *Config {
	log.Println("gather config")

	instance := &Config{}
	if err := cleanenv.ReadEnv(instance); err != nil {
		log.Fatalf("failed to gather config: %s", err.Error())
	}

	return instance
}

var Get = sync.OnceValue(_get)
