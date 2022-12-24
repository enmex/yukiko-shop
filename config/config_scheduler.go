package config

import (
	"os"
	"time"
	"yukiko-shop/pkg/scheduler"
)

func NewConfigScheduler() (*scheduler.Config, error) {
	frequency, err := time.ParseDuration(os.Getenv("SCHEDULER_FREQUENCY"))
	if err != nil {
		return nil, err
	}

	return &scheduler.Config{
		Frequency: frequency,
	}, nil
}
