package configuration

import (
	"time"

	"github.com/Symantec/Dominator/lib/log"
)

type Configuration struct {
}

func Watch(configUrl string, checkInterval time.Duration,
	logger log.DebugLogger) (<-chan *Configuration, error) {
	ch := make(chan *Configuration, 1)
	return ch, nil
}