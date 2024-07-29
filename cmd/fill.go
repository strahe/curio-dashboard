package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/strahe/curio-dashboard/aggregator"
	"github.com/strahe/curio-dashboard/aggregator/jobs"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph/loaders"
	"github.com/urfave/cli/v2"
)

var fillCmd = &cli.Command{
	Name:  "fill",
	Usage: "fill gaps in the aggregator database for a given range and set of jobs",
	Flags: []cli.Flag{
		&cli.TimestampFlag{
			Name:     "start",
			Usage:    "start time for the fill, e.g., '2006-01-02 15:04:05'",
			Layout:   time.DateTime,
			Required: true,
		},
		&cli.TimestampFlag{
			Name:        "end",
			Usage:       "end time for the fill, e.g., '2006-01-02 15:04:05'",
			Layout:      time.DateTime,
			DefaultText: time.Now().Format(time.DateTime),
			Value:       cli.NewTimestamp(time.Now()),
		},
		&cli.StringSliceFlag{
			Name:  "jobs",
			Usage: "list of jobs to run, all jobs will run if not specified",
		},
		harmonydbFlag,
		appdbFlag,
	},
	Action: func(cctx *cli.Context) error {
		startTime := cctx.Timestamp("start")
		endTime := cctx.Timestamp("end")

		harmonyDB, err := db.NewHarmonyDB(cctx.Context, cctx.String("harmonydb-url"))
		if err != nil {
			return fmt.Errorf("failed to connect to harmonydb: %w", err)
		}
		defer harmonyDB.Close()

		appDB, err := db.NewAppDb(cctx.String("appdb-url"))
		if err != nil {
			return fmt.Errorf("failed to connect to appdb: %w", err)
		}

		if os.Getenv("FULLNODE_API_INFO") == "" {
			return fmt.Errorf("FULLNODE_API_INFO not set")
		}
		apiInfo := strings.Split(os.Getenv("FULLNODE_API_INFO"), ",")
		fullNode, closer, err := getFullNodeAPIV1(cctx, apiInfo)
		if err != nil {
			return fmt.Errorf("failed to get full node API: %w", err)
		}
		defer closer()

		var runJobs []jobs.BackFillingJob
		if cctx.IsSet("jobs") { // nolint:revive,empty-block
			// todo: parse jobs
		} else {
			allJobs := aggregator.AllJobs(harmonyDB, appDB, fullNode, loaders.NewLoader(harmonyDB, appDB, 1000))
			runJobs = aggregator.AllBackFillingJobs(allJobs)
		}

		parser := cron.NewParser(cron.Second | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
		for _, job := range runJobs {
			jobSchedule, err := parser.Parse(job.Spec())
			if err != nil {
				return fmt.Errorf("parsing job schedule: %w", err)
			}
			// todo: maybe parallelize this
			for t := jobSchedule.Next(*startTime); t.Before(*endTime); t = jobSchedule.Next(t) {
				log.Infof("Running job %s at %s", job.Name(), t)
				job.RunWith(t)
			}
		}
		return nil
	},
}
