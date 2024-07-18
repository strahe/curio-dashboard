package jobs

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
	"time"
)

var log = logging.Logger("jobs")

type Job interface {
	cron.Job
	RunWith(time time.Time) //
	Spec() string           // Spec returns the cron expression that determines the job schedule.
	Name() string
}
