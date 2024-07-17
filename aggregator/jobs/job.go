package jobs

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
)

var log = logging.Logger("jobs")

type Job interface {
	cron.Job
	Spec() string // Spec returns the cron expression that determines the job schedule.
	Name() string
}
