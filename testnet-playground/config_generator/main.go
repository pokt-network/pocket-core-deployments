package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	apps "github.com/pokt-network/pocket-core/x/apps"
	appsTypes "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/pokt-network/pocket-core/x/nodes"
	nodesTypes "github.com/pokt-network/pocket-core/x/nodes/types"
	pocket "github.com/pokt-network/pocket-core/x/pocketcore"
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"github.com/pokt-network/posmint/crypto"
	sdk "github.com/pokt-network/posmint/types"
	"github.com/pokt-network/posmint/types/module"
	"github.com/pokt-network/posmint/x/auth"
	"github.com/pokt-network/posmint/x/bank"
	"github.com/pokt-network/posmint/x/params"
	"github.com/pokt-network/posmint/x/supply"
	tmTypes "github.com/tendermint/tendermint/types"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	var urls []string
	numberOfNodes, err := strconv.Atoi(os.Args[1])
	urls = append(urls, os.Args[2])
	urls = append(urls, os.Args[3])
	if err != nil {
		panic(err)
	}
	k := keys(numberOfNodes)
	genesis(k)
	chains(urls)
}

type JSONKeys struct {
	Priv string `json:"private_key"`
	Pub  string `json:"public_key"`
	Addr string `json:"address"`
}

func keys(n int) []JSONKeys {
	var res []JSONKeys
	for i := 0; i < n; i++ {
		pk := crypto.GenerateEd25519PrivKey()
		res = append(res, JSONKeys{
			Priv: pk.RawString(),
			Pub:  pk.PublicKey().RawString(),
			Addr: pk.PublicKey().Address().String(),
		})
	}
	bz, _ := json.MarshalIndent(res, "", "  ")
	err := ioutil.WriteFile("keys.json", bz, 0644)
	if err != nil {
		panic(err)
	}
	return res
}

func genesis(keys []JSONKeys) {
	j, er := types.ModuleCdc.MarshalJSONIndent(tmTypes.GenesisDoc{
		GenesisTime: time.Now(),
		ChainID:     "pocket-test",
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
		Validators: nil,
		AppHash:    nil,
		AppState:   newDefaultGenesisState(keys),
	}, "", "  ")
	if er != nil {
		panic(er)
	}
	er = ioutil.WriteFile("genesis.json", j, 0644)
	if er != nil {
		panic(er)
	}
}

func newDefaultGenesisState(keys []JSONKeys) []byte {
	defaultGenesis := module.NewBasicManager(
		apps.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		nodes.AppModuleBasic{},
		supply.AppModuleBasic{},
		pocket.AppModuleBasic{},
	).DefaultGenesis()
	ethereum, err := types.NonNativeChain{
		Ticker:  "eth",
		Netid:   "4",
		Version: "v1.9.9",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
	bitcoin, err := types.NonNativeChain{
		Ticker:  "btc",
		Netid:   "1",
		Version: "0.19.0.1",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
	serviceURLPrefix := "http://www.pocket-core-testnet"
	rawPOS := defaultGenesis[nodesTypes.ModuleName]
	var posGenesisState nodesTypes.GenesisState
	types.ModuleCdc.MustUnmarshalJSON(rawPOS, &posGenesisState)
	for i, jk := range keys {
		addr, err := hex.DecodeString(jk.Addr)
		if err != nil {
			panic(err)
		}
		pk, err := crypto.NewPublicKey(jk.Pub)
		if err != nil {
			panic(err)
		}
		posGenesisState.Validators = append(posGenesisState.Validators,
			nodesTypes.Validator{Address: addr,
				PublicKey:    pk,
				Status:       sdk.Bonded,
				Chains:       []string{ethereum, bitcoin},
				ServiceURL:   serviceURLPrefix + strconv.Itoa(i) + ":8081",
				StakedTokens: sdk.NewInt(1000000000000)})
	}
	res := types.ModuleCdc.MustMarshalJSON(posGenesisState)
	defaultGenesis[nodesTypes.ModuleName] = res
	// setup a single application
	var appsGenesisState appsTypes.GenesisState
	rawApps := defaultGenesis[appsTypes.ModuleName]
	types.ModuleCdc.MustUnmarshalJSON(rawApps, &appsGenesisState)
	coinBase := posGenesisState.Validators[0]
	appsGenesisState.Applications = append(appsGenesisState.Applications, appsTypes.Application{
		Address:      coinBase.Address,
		PublicKey:    coinBase.PublicKey,
		Jailed:       false,
		Status:       sdk.Bonded,
		Chains:       []string{ethereum, bitcoin},
		StakedTokens: sdk.NewInt(1000000000000000),
	})
	res = appsTypes.ModuleCdc.MustMarshalJSON(appsGenesisState)
	defaultGenesis[appsTypes.ModuleName] = res
	// setup a single account
	var accsGenesisState auth.GenesisState
	rawAccs := defaultGenesis[auth.ModuleName]
	types.ModuleCdc.MustUnmarshalJSON(rawAccs, &accsGenesisState)
	accsGenesisState.Accounts = append(accsGenesisState.Accounts, &auth.BaseAccount{
		Address: coinBase.Address,
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000000000000))),
		PubKey:  coinBase.PublicKey,
	})
	res = auth.ModuleCdc.MustMarshalJSON(accsGenesisState)
	defaultGenesis[auth.ModuleName] = res
	j, er := types.ModuleCdc.MarshalJSONIndent(defaultGenesis, "", "    ")
	if er != nil {
		panic(er)
	}
	fmt.Println(coinBase)
	return j
}

func chains(urls []string) {
	ethereum, err := types.NonNativeChain{
		Ticker:  "eth",
		Netid:   "4",
		Version: "v1.9.9",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
	bitcoin, err := types.NonNativeChain{
		Ticker:  "btc",
		Netid:   "1",
		Version: "0.19.0.1",
		Client:  "",
		Inter:   "",
	}.HashString()
	if err != nil {
		panic(err)
	}
	hbs := types.HostedBlockchains{M: map[string]types.HostedBlockchain{
		ethereum: {
			Hash: ethereum,
			URL:  urls[0],
		},
		bitcoin: {
			Hash: bitcoin,
			URL:  urls[1],
		},
	}}
	j, er := types.ModuleCdc.MarshalJSONIndent(hbs.M, "", "    ")
	if er != nil {
		panic(er)
	}
	er = ioutil.WriteFile("chains.json", j, 0644)
	if er != nil {
		panic(er)
	}
}
