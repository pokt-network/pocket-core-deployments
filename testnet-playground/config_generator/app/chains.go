package app

import (
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"io/ioutil"
	"path/filepath"
)

var (
	fs = string(filepath.Separator)
)

func GenChains(homeDir, ethereumURL, bitcoinURL string) {
	hbs := types.HostedBlockchains{M: map[string]types.HostedBlockchain{
		ethereum: {
			Hash: ethereum,
			URL:  ethereumURL,
		},
		bitcoin: {
			Hash: bitcoin,
			URL:  bitcoinURL,
		},
	}}
	chainsJSON, er := types.ModuleCdc.MarshalJSONIndent(hbs.M, "", "    ")
	if er != nil {
		panic(er)
	}
	er = ioutil.WriteFile(homeDir+fs+"chains.json", chainsJSON, 0644)
	if er != nil {
		panic(er)
	}
}
