package app

import (
	"encoding/hex"
	"encoding/json"
	apps "github.com/pokt-network/pocket-core/x/apps"
	appsTypes "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/pokt-network/pocket-core/x/nodes"
	nodesTypes "github.com/pokt-network/pocket-core/x/nodes/types"
	pocket "github.com/pokt-network/pocket-core/x/pocketcore"
	pocketTypes "github.com/pokt-network/pocket-core/x/pocketcore/types"
	"github.com/pokt-network/posmint/crypto"
	sdk "github.com/pokt-network/posmint/types"
	"github.com/pokt-network/posmint/types/module"
	"github.com/pokt-network/posmint/x/auth"
	"github.com/pokt-network/posmint/x/bank"
	"github.com/pokt-network/posmint/x/params"
	"github.com/pokt-network/posmint/x/supply"
	tmTypes "github.com/tendermint/tendermint/types"
	"io/ioutil"
	"strconv"
	"time"
)

var (
	ethereum, bitcoin string
)

func init() {
	// setup two non native chains
	var err error
	ethereum, err = pocketTypes.NonNativeChain{
		Ticker:  "eth",
		Netid:   "4",
		Version: "v1.9.9",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
	bitcoin, err = pocketTypes.NonNativeChain{
		Ticker:  "btc",
		Netid:   "1",
		Version: "0.19.0.1",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
}

func GenGenesis(homeDir string, keys KeysFile, minutesTillGenesisStart int) {
	j, er := pocketTypes.ModuleCdc.MarshalJSONIndent(tmTypes.GenesisDoc{
		GenesisTime: time.Now().Add(time.Minute * time.Duration(minutesTillGenesisStart)),
		ChainID:     "pocket-testet-playground",
		ConsensusParams: &tmTypes.ConsensusParams{
			Block: tmTypes.BlockParams{
				MaxBytes:   15000,
				MaxGas:     -1,
				TimeIotaMs: 1,
			},
			Evidence: tmTypes.EvidenceParams{
				MaxAge: 1000000,
			},
			Validator: tmTypes.ValidatorParams{
				PubKeyTypes: []string{"ed25519"},
			},
		},
		AppState: newAppState(keys),
	}, "", "  ")
	if er != nil {
		panic(er)
	}
	er = ioutil.WriteFile(homeDir+fs+"genesis.json", j, 0644)
	if er != nil {
		panic(er)
	}
}

func newAppState(keys KeysFile) []byte {
	// setup the default geneis to start
	defaultGenesis := module.NewBasicManager(
		apps.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		nodes.AppModuleBasic{},
		supply.AppModuleBasic{},
		pocket.AppModuleBasic{},
	).DefaultGenesis()
	setupNodeGenesis(defaultGenesis, keys)
	setupAppGenesis(defaultGenesis, keys)
	setupAccGenesis(defaultGenesis, keys)
	setupPocketGenesis(defaultGenesis)
	genesisJSON, er := pocketTypes.ModuleCdc.MarshalJSONIndent(defaultGenesis, "", "    ")
	if er != nil {
		panic(er)
	}
	return genesisJSON
}

func setupAccGenesis(defaultGenesis map[string]json.RawMessage, keys KeysFile) {
	// setup accounts
	var accGenesisObj auth.GenesisState
	rawAccGenesis := defaultGenesis[auth.ModuleName]
	pocketTypes.ModuleCdc.MustUnmarshalJSON(rawAccGenesis, &accGenesisObj)
	for _, app := range keys.AppKeys {
		addr, err := hex.DecodeString(app.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(app.Pub)
		if err != nil {
			panic(err)
		}
		accGenesisObj.Accounts = append(accGenesisObj.Accounts, &auth.BaseAccount{
			Address: addr,
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultStakeDenom, sdk.NewInt(1000000000))),
			PubKey:  pk,
		})
	}
	for _, app := range keys.NodeKeys {
		addr, err := hex.DecodeString(app.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(app.Pub)
		if err != nil {
			panic(err)
		}
		accGenesisObj.Accounts = append(accGenesisObj.Accounts, &auth.BaseAccount{
			Address: addr,
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultStakeDenom, sdk.NewInt(1000000000))),
			PubKey:  pk,
		})
	}
	for _, app := range keys.AccKeys {
		addr, err := hex.DecodeString(app.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(app.Pub)
		if err != nil {
			panic(err)
		}
		accGenesisObj.Accounts = append(accGenesisObj.Accounts, &auth.BaseAccount{
			Address: addr,
			Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultStakeDenom, sdk.NewInt(1000000000))),
			PubKey:  pk,
		})
	}
	res := auth.ModuleCdc.MustMarshalJSON(accGenesisObj)
	defaultGenesis[auth.ModuleName] = res
}

func setupAppGenesis(defaultGenesis map[string]json.RawMessage, keys KeysFile) {
	var appsGenesisObj appsTypes.GenesisState
	rawAppsGenesis := defaultGenesis[appsTypes.ModuleName]
	pocketTypes.ModuleCdc.MustUnmarshalJSON(rawAppsGenesis, &appsGenesisObj)
	for _, app := range keys.AppKeys {
		addr, err := hex.DecodeString(app.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(app.Pub)
		if err != nil {
			panic(err)
		}
		appsGenesisObj.Applications = append(appsGenesisObj.Applications, appsTypes.Application{
			Address:      addr,
			PublicKey:    pk,
			Jailed:       false,
			Status:       sdk.Staked,
			Chains:       []string{ethereum, bitcoin},
			StakedTokens: sdk.NewInt(1000000000),
		})
	}
	res := appsTypes.ModuleCdc.MustMarshalJSON(appsGenesisObj)
	defaultGenesis[appsTypes.ModuleName] = res
}

func setupNodeGenesis(defaultGenesis map[string]json.RawMessage, keys KeysFile) {
	// setup the service url prefix
	serviceURLPrefix := "http://www.pocket-core-testnet"
	rawNodesGenesis := defaultGenesis[nodesTypes.ModuleName]
	var nodesGenesisObj nodesTypes.GenesisState
	pocketTypes.ModuleCdc.MustUnmarshalJSON(rawNodesGenesis, &nodesGenesisObj)
	for i, node := range keys.NodeKeys {
		addr, err := hex.DecodeString(node.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(node.Pub)
		if err != nil {
			panic(err)
		}
		nodesGenesisObj.Validators = append(nodesGenesisObj.Validators,
			nodesTypes.Validator{Address: addr,
				PublicKey:    pk,
				Status:       sdk.Staked,
				Chains:       []string{ethereum, bitcoin},
				ServiceURL:   serviceURLPrefix + strconv.Itoa(i) + ":8081",
				StakedTokens: sdk.NewInt(1000000000)})
	}
	res := pocketTypes.ModuleCdc.MustMarshalJSON(nodesGenesisObj)
	defaultGenesis[nodesTypes.ModuleName] = res
}

func setupPocketGenesis(defaultGenesis map[string]json.RawMessage) {
	// setup the service url prefix
	rawPocketGenesis := defaultGenesis[pocketTypes.ModuleName]
	var pocketGenesisObj pocketTypes.GenesisState
	pocketTypes.ModuleCdc.MustUnmarshalJSON(rawPocketGenesis, &pocketGenesisObj)
	pocketGenesisObj.Params.SupportedBlockchains = append(pocketGenesisObj.Params.SupportedBlockchains, ethereum)
	pocketGenesisObj.Params.SupportedBlockchains = append(pocketGenesisObj.Params.SupportedBlockchains, bitcoin)
	res := pocketTypes.ModuleCdc.MustMarshalJSON(pocketGenesisObj)
	defaultGenesis[pocketTypes.ModuleName] = res
}
