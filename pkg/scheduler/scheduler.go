package scheduler

import (
	"context"
	"time"
)

type Scheduler struct {
	callback func(ctx context.Context) error
	errChan  chan error
	ticker   time.Ticker
}

func NewScheduler(cfg *Config, callback func(ctx context.Context) error) *Scheduler {
	return &Scheduler{
		ticker:   *time.NewTicker(cfg.Frequency),
		callback: callback,
		errChan:  make(chan error),
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	go func() {
		for {
			<-s.ticker.C
			if err := s.callback(ctx); err != nil {
				s.errChan <- err
			}
		}
	}()
}

func (s *Scheduler) Error() chan error {
	return s.errChan
}
