package main

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/alecthomas/kong"
	"github.com/dontpanicdao/caigo/jsonrpc"
	"github.com/dontpanicdao/caigo/types"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	indexer "github.com/tarrencev/starknet-indexer"
	_ "github.com/tarrencev/starknet-indexer/ent/runtime"
)

func main() {
	var cli struct {
		Addr  string `name:"address" default:":8081" help:"Address to listen on."`
		Debug bool   `name:"debug" help:"Enable debugging mode."`
	}
	kong.Parse(&cli)

	drv, err := sql.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
	)
	if err != nil {
		log.Fatal().Err(err).Msg("opening ent client")
	}

	provider, err := jsonrpc.DialContext(context.Background(), "http://localhost:9545")
	if err != nil {
		log.Fatal().Err(err).Msg("Dialing provider")
	}

	indexer.New(cli.Addr, drv, provider, indexer.Config{
		Interval: 2 * time.Second,
		Contracts: []indexer.Contract{{
			Address:    "0x",
			StartBlock: 1000,
			Handler: func(types.Transaction) error {
				// handle transaction
				return nil
			},
		}},
	})
}
