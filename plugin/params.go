package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/coreth/plugin/evm"
)

const (
	name = "coreth"
)

var (
	cliConfig evm.CommandLineConfig
	errs      wrappers.Errs
)

func init() {
	errs := wrappers.Errs{}
	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	config := fs.String("config", "default", "Pass in CLI Config to set runtime attributes for Coreth")

	if err := fs.Parse(os.Args[1:]); err != nil {
		errs.Add(err)
		return
	}

	if *config == "default" {
		cliConfig.EthAPIEnabled = true
		cliConfig.PersonalAPIEnabled = true
		cliConfig.TxPoolAPIEnabled = true
		cliConfig.RPCGasCap = 2500000000 // 25000000 x 100
		cliConfig.RPCTxFeeCap = 100      // 100 AVAX
	} else {
		// TODO only overwrite values that were explicitly set
		errs.Add(json.Unmarshal([]byte(*config), &cliConfig))
	}
}
