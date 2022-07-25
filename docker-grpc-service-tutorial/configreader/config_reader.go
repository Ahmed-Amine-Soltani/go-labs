package configreader

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// These global variables makes it easy
// to mock these dependencies
// in unit tests.
var (
	godotenvLoad     = godotenv.Load
	envconfigProcess = envconfig.Process
)

// GoDotEnv is an interface that defines
// the functions we use from godotenv package.
// It enables mocking this dependency in unit testing.
type GoDotEnv interface {
	Load(filenames ...string) (err error)
}

// EnvConfig is an interface that defines
// the functions we use from envconfig package.
// It enables mocking this dependency in unit testing.
type EnvConfig interface {
	Process(prefix string, spec interface{}) error
}

// Config holds configuration data.
type Config struct {
	PoetrydbBaseUrl     string `envconfig:"POETRYDB_BASE_URL" required:"true"`
	PoetrydbHttpTimeout int    `envconfig:"POETRYDB_HTTP_TIMEOUT" required:"true"`
}

// ReadEnv reads envionment variables into Config struct.
func ReadEnv() (*Config, error) {
	err := godotenvLoad("configreader/config.env")
	if err != nil {
		return nil, errors.Wrap(err, "reading .env file")
	}
	var config Config
	err = envconfigProcess("", &config)
	if err != nil {
		return nil, errors.Wrap(err, "processing env vars")
	}
	return &config, nil
}
