package backend

import (
	"github.com/Dataman-Cloud/swan/src/manager/sched/scheduler"
	. "github.com/Dataman-Cloud/swan/src/manager/store"
)

type Backend struct {
	sched *scheduler.Scheduler
	store Store
}

func NewBackend(sched *scheduler.Scheduler, store Store) *Backend {
	return &Backend{
		sched: sched,
		store: store,
	}
}

func (b *Backend) ClusterId() string {
	return b.sched.ClusterId
}
