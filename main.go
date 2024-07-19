package main

import (
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/strahe/curio-dashboard/graph/loaders"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v1api"
	"github.com/filecoin-project/lotus/build"
	lcli "github.com/filecoin-project/lotus/cli"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/gorilla/websocket"
	logging "github.com/ipfs/go-log/v2"
	"github.com/robfig/cron/v3"
	"github.com/rs/cors"
	"github.com/strahe/curio-dashboard/aggregator"
	"github.com/strahe/curio-dashboard/aggregator/jobs"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph"
	cachecontrol "github.com/strahe/curio-dashboard/graph/cache_control"
	"github.com/strahe/curio-dashboard/graph/resolvers"
	"github.com/strahe/curio-dashboard/web"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var log = logging.Logger("main")

func main() {
	app := &cli.App{
		Name:                 "curio-dashboard",
		Usage:                "A dashboard for the Curio",
		Version:              "0.0.1", // todo:
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
				Value: false,
			},
		},
		Before: func(c *cli.Context) error {
			if c.Bool("debug") {
				return logging.SetLogLevel("*", "DEBUG")
			} else {
				return logging.SetLogLevel("*", "INFO")
			}
		},
		After: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			runCmd,
			fillCmd,
		},
	}
	app.Setup()
	lcli.RunApp(app)
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "run the Curio dashboard server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "listen",
			Usage: "host address and port will listen on",
			Value: "127.0.0.1:9091",
		},
		harmonydbFlag,
		appdbFlag,
		cliutil.FlagVeryVerbose,
	},
	Action: func(cctx *cli.Context) error {
		harmonyDB, err := db.NewHarmonyDB(cctx.Context, cctx.String("harmonydb-url"))
		if err != nil {
			return fmt.Errorf("failed to connect to harmonydb: %w", err)
		}
		defer harmonyDB.Close()
		if os.Getenv("FULLNODE_API_INFO") == "" {
			return fmt.Errorf("FULLNODE_API_INFO not set")
		}
		apiInfo := strings.Split(os.Getenv("FULLNODE_API_INFO"), ",")
		fullNode, closer, err := getFullNodeAPIV1(cctx, apiInfo)
		if err != nil {
			return fmt.Errorf("failed to get full node API: %w", err)
		}
		defer closer()

		// get network name
		ntn, err := fullNode.StateNetworkName(cctx.Context)
		if err != nil {
			return fmt.Errorf("failed to get network name: %w", err)
		}
		log.Infof("Using network: %s", ntn)
		if err := build.UseNetworkBundle(string(ntn)); err != nil {
			return fmt.Errorf("failed to use network bundle: %w", err)
		}
		if ntn == "calibrationnet" {
			address.CurrentNetwork = address.Testnet
		} else {
			address.CurrentNetwork = address.Mainnet
		}

		appDB, err := db.NewAppDb(cctx.Context, cctx.String("appdb-url"))
		if err != nil {
			return fmt.Errorf("failed to connect to appdb: %w", err)
		}

		agg, err := aggregator.NewManager(fullNode, harmonyDB, appDB)
		if err != nil {
			return fmt.Errorf("failed to create aggregator manager: %w", err)
		}
		agg.Start()
		defer agg.Stop()

		router := http.NewServeMux()
		var fn func(r *http.Request) bool
		if cctx.Bool("debug") {
			fn = func(r *http.Request) bool {
				return true
			}
		}

		var srv = handler.New(
			graph.NewExecutableSchema(graph.Config{Resolvers: resolvers.NewResolver(harmonyDB, appDB, fullNode)}))
		srv.AddTransport(transport.Websocket{
			KeepAlivePingInterval: 10 * time.Second,
			Upgrader: websocket.Upgrader{
				CheckOrigin: fn,
			},
		})
		srv.AddTransport(transport.Options{})
		srv.AddTransport(transport.GET{})
		srv.AddTransport(transport.POST{})
		srv.AddTransport(transport.MultipartForm{})
		srv.SetQueryCache(lru.New(1000))
		srv.Use(extension.Introspection{})
		srv.Use(extension.AutomaticPersistedQuery{
			Cache: lru.New(100),
		})
		srv.Use(cachecontrol.Extension{})

		assets, _ := web.Assets()
		fs := http.FileServer(http.FS(assets))
		router.Handle("/", http.StripPrefix("/", fs))
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
		if cctx.Bool("debug") {
			router.Handle("/graphql", cors.AllowAll().Handler(srv))
		} else {
			router.Handle("/graphql", srv)
		}
		listen := cctx.String("listen")

		log.Infof("connect to %s for GraphQL playground", listen)
		log.Fatal(http.ListenAndServe(listen, router))
		return nil
	},
}

func getFullNodeAPIV1(ctx *cli.Context, aiCfg []string, opts ...cliutil.GetFullNodeOption) (v1api.FullNode, jsonrpc.ClientCloser, error) {

	var options cliutil.GetFullNodeOptions
	for _, opt := range opts {
		opt(&options)
	}

	var rpcOpts []jsonrpc.Option
	if options.EthSubHandler != nil {
		rpcOpts = append(rpcOpts, jsonrpc.WithClientHandler("Filecoin", options.EthSubHandler), jsonrpc.WithClientHandlerAlias("eth_subscription", "Filecoin.EthSubscription"))
	}

	var httpHeads []httpHead
	version := "v1"
	{
		if len(aiCfg) == 0 {
			return nil, nil, xerrors.Errorf("could not get API info: none configured. \nConsider getting base.toml with './curio config get base >/tmp/base.toml' \nthen adding   \n[APIs] \n ChainApiInfo = [\" result_from lotus auth api-info --perm=admin \"]\n  and updating it with './curio config set /tmp/base.toml'")
		}
		for _, i := range aiCfg {
			ai := cliutil.ParseApiInfo(i)
			addr, err := ai.DialArgs(version)
			if err != nil {
				return nil, nil, xerrors.Errorf("could not get DialArgs: %s", err)
			}
			httpHeads = append(httpHeads, httpHead{addr: addr, header: ai.AuthHeader()})
		}
	}

	if cliutil.IsVeryVerbose {
		_, _ = fmt.Fprintln(ctx.App.Writer, "using full node API v1 endpoint:", httpHeads[0].addr)
	}

	var fullNodes []api.FullNode
	var closers []jsonrpc.ClientCloser

	for _, head := range httpHeads {
		v1, closer, err := client.NewFullNodeRPCV1(ctx.Context, head.addr, head.header, rpcOpts...)
		if err != nil {
			log.Warnf("Not able to establish connection to node with addr: %s, Reason: %s", head.addr, err.Error())
			continue
		}
		fullNodes = append(fullNodes, v1)
		closers = append(closers, closer)
	}

	// When running in cluster mode and trying to establish connections to multiple nodes, fail
	// if less than 2 lotus nodes are actually running
	if len(httpHeads) > 1 && len(fullNodes) < 2 {
		return nil, nil, xerrors.Errorf("Not able to establish connection to more than a single node")
	}

	finalCloser := func() {
		for _, c := range closers {
			c()
		}
	}

	var v1API api.FullNodeStruct
	cliutil.FullNodeProxy(fullNodes, &v1API)

	v, err := v1API.Version(ctx.Context)
	if err != nil {
		return nil, nil, err
	}
	if !v.APIVersion.EqMajorMinor(api.FullAPIVersion1) {
		return nil, nil, xerrors.Errorf("Remote API version didn't match (expected %s, remote %s)", api.FullAPIVersion1, v.APIVersion)
	}
	return &v1API, finalCloser, nil
}

type httpHead struct {
	addr   string
	header http.Header
}

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

		appDB, err := db.NewAppDb(cctx.Context, cctx.String("appdb-url"))
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

		var runJobs []jobs.Job
		if cctx.IsSet("jobs") {
			// todo: parse jobs
		} else {
			runJobs = aggregator.AllOnlyRealtimeJobs(appDB, fullNode, loaders.NewLoader(harmonyDB, appDB, 1000))
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

var harmonydbFlag = &cli.StringFlag{
	Name:     "harmonydb-url",
	EnvVars:  []string{"CURIO_HARMONYDB_URL"},
	Usage:    "URL to connect to harmonydb, e.g. postgres://username:password@localhost:5433/database_name?search_path=curio",
	Required: true,
}

var appdbFlag = &cli.StringFlag{
	Name:    "appdb-url",
	EnvVars: []string{"CURIO_APPDB_URL"},
	Value:   "sqlite3://app.sqlite",
	Usage:   "URL to connect to appdb, e.g. postgres://username:password@localhost:5433/database_name?search_path=dashboard",
}
