package app

import (
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var (
	fs = string(filepath.Separator)
)

func GenChains(homeDir, ethereumURL, bitcoinURL string) string {
	hbs := types.HostedBlockchains{M: map[string]types.HostedBlockchain{
		ethereum: {
			Hash: ethereum,
			URL:  strings.TrimSuffix(ethereumURL, "\n"),
		},
		bitcoin: {
			Hash: bitcoin,
			URL:  strings.TrimSuffix(bitcoinURL, "\n"),
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
	return string(chainsJSON)
}
