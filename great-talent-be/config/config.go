package config

import (
	"github.com/joho/godotenv"
	"great-talent-be/exception"
	"os"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filename ...string) Config {
	err := godotenv.Load(filename...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
