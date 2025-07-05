package model

import (
	"context"
	"time"
)

type Task struct {
	Status    string
	StartTime time.Time
	Duration  time.Duration
	Ctx       context.Context
	Cancel    context.CancelFunc
}
