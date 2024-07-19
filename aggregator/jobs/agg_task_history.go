package jobs

import (
	"time"

	"github.com/strahe/curio-dashboard/aggregator/query"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AggTaskHistory struct {
	query *query.Query
	db    *db.HarmonyDB
	appDB *gorm.DB
}

func NewAggTaskHistory(query *query.Query, db *db.HarmonyDB, appDB *gorm.DB) *AggTaskHistory {
	return &AggTaskHistory{
		query: query,
		db:    db,
		appDB: appDB,
	}
}

func (j *AggTaskHistory) Spec() string {
	return "@hourly"
}

func (j *AggTaskHistory) Name() string {
	return "AggTaskHistory"
}

func (j *AggTaskHistory) RunWith(t time.Time) {
	start := time.Now()
	defer func() {
		log.Infof("%s finished in %s", j.Name(), time.Since(start))
	}()

	currentHour := t.Truncate(time.Hour)
	ags, err := j.query.AggregateTaskHistory(context.TODO(), currentHour.Add(-time.Hour), currentHour)
	if err != nil {
		log.Errorf("Failed to aggregate task history: %s", err)
		return
	}

	var records []*model.TaskHistoryAggregates

	for _, ag := range ags {
		records = append(records, &model.TaskHistoryAggregates{
			Time:    ag.Hour,
			Task:    ag.Name,
			Total:   ag.Total,
			Success: ag.Success,
			Failure: ag.Failure,
		})
	}
	if len(records) == 0 {
		return
	}
	j.appDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Save(records)
}

func (j *AggTaskHistory) Run() {
	t := time.Now()
	j.RunWith(t)
}
