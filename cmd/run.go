package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/build"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/strahe/curio-dashboard/aggregator"
	"github.com/strahe/curio-dashboard/db"
	"github.com/strahe/curio-dashboard/graph"
	cachecontrol "github.com/strahe/curio-dashboard/graph/cachecontrol"
	"github.com/strahe/curio-dashboard/graph/resolvers"
	"github.com/strahe/curio-dashboard/web"
	"github.com/urfave/cli/v2"
)

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
		} else if strings.HasPrefix(string(ntn), "localnet") {
			address.CurrentNetwork = address.Testnet
			if err := build.UseNetworkBundle("devnet"); err != nil {
				return fmt.Errorf("failed to use network bundle: %w", err)
			}
		} else {
			address.CurrentNetwork = address.Mainnet
		}

		appDB, err := db.NewAppDb(cctx.String("appdb-url"))
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
			fn = func(*http.Request) bool {
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

		assets, _ := web.Assets() // nolint:errcheck
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
		log.Fatal(http.ListenAndServe(listen, router)) // nolint: gosec
		return nil
	},
}
