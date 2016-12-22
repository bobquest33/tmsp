package main

import (
	"flag"

	. "github.com/tendermint/go-common"
	"github.com/tendermint/abci/server"
	"github.com/tendermint/abci/types"
)

func main() {

	addrPtr := flag.String("addr", "tcp://0.0.0.0:46658", "Listen address")
	abciPtr := flag.String("abci", "socket", "socket | grpc")
	flag.Parse()

	// Start the listener
	_, err := server.NewServer(*addrPtr, *abciPtr, NewChainAwareApplication())
	if err != nil {
		Exit(err.Error())
	}

	// Wait forever
	TrapSignal(func() {
		// Cleanup
	})

}

type ChainAwareApplication struct {
	beginCount int
	endCount   int
}

func NewChainAwareApplication() *ChainAwareApplication {
	return &ChainAwareApplication{}
}

func (app *ChainAwareApplication) Info() (string, *types.StateInfo, *types.ConfigInfo) {
	return "nil", nil, nil, nil
}

func (app *ChainAwareApplication) SetOption(key string, value string) (log string) {
	return ""
}

func (app *ChainAwareApplication) DeliverTx(tx []byte) types.Result {
	return types.NewResultOK(nil, "")
}

func (app *ChainAwareApplication) CheckTx(tx []byte) types.Result {
	return types.NewResultOK(nil, "")
}

func (app *ChainAwareApplication) Commit() types.Result {
	return types.NewResultOK([]byte("nil"), "")
}

func (app *ChainAwareApplication) Query(query []byte) types.Result {
	return types.NewResultOK([]byte(Fmt("%d,%d", app.beginCount, app.endCount)), "")
}

func (app *ChainAwareApplication) BeginBlock(hash []byte, header *types.Header) {
	app.beginCount += 1
	return
}

func (app *ChainAwareApplication) EndBlock(height uint64) []*types.Validator {
	app.endCount += 1
	return nil
}

func (app *ChainAwareApplication) InitChain(vals []*types.Validator) {
	return
}
