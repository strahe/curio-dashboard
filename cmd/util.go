package main

import (
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/api/client"
	"github.com/filecoin-project/lotus/api/v1api"
	cliutil "github.com/filecoin-project/lotus/cli/util"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

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
