package app

import (
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"io/ioutil"
	"path/filepath"
)

var (
	fs            = string(filepath.Separator)
	defaultChains = types.HostedBlockchains{M: map[string]types.HostedBlockchain{
		"8ef9a7c67f6f8ad14f82c1f340963951245f912f037a7087f3f2d2f9f9ee38a8": {
			Hash: "8ef9a7c67f6f8ad14f82c1f340963951245f912f037a7087f3f2d2f9f9ee38a8",
			URL:  "http://aion-mainnet.prod.svc.cluster.local:8545",
		},
		"0070eebec778ea95ef9c75551888971c27cce222e00b2f3f79168078b8a77ff9": {
			Hash: "0070eebec778ea95ef9c75551888971c27cce222e00b2f3f79168078b8a77ff9",
			URL:  "http://daps-mainnet.prod.svc.cluster.local:53573",
		},
		"4ae7539e01ad2c42528b6a697f118a3535e404fe65999b2c6fee506465390367": {
			Hash: "4ae7539e01ad2c42528b6a697f118a3535e404fe65999b2c6fee506465390367",
			URL:  "http://geth-goerli.prod.svc.cluster.local:8545",
		},
		"0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b": {
			Hash: "0de3141aec1e69aea9d45d9156269b81a3ab4ead314fbf45a8007063879e743b",
			URL:  "http://geth-mainnet.prod.svc.cluster.local:8545",
		},
		"8cf7f8799c5b30d36c86d18f0f4ca041cf1803e0414ed9e9fd3a19ba2f0938ff": {
			Hash: "8cf7f8799c5b30d36c86d18f0f4ca041cf1803e0414ed9e9fd3a19ba2f0938ff",
			URL:  "http://geth-rinkeby.prod.svc.cluster.local:8545",
		},
		"10d1290eee169e3970afb106fe5417a11b81676ce1e2119a0292df29f0445d30": {
			Hash: "10d1290eee169e3970afb106fe5417a11b81676ce1e2119a0292df29f0445d30",
			URL:  "http://geth-ropsten-2.prod.svc.cluster.local:8545",
		},
		"d9d669583c2d2a88e54c0120be6f8195b2575192f178f925099813ff9095d139": {
			Hash: "d9d669583c2d2a88e54c0120be6f8195b2575192f178f925099813ff9095d139",
			URL:  "",
		},
		"d9d77bce50d80e70026bd240fb0759f08aab7aee63d0a6d98c545f2b5ae0a0b8": {
			Hash: "d9d77bce50d80e70026bd240fb0759f08aab7aee63d0a6d98c545f2b5ae0a0b8",
			URL:  "",
		},
		"dcc98e38e1edb55a97265efca6c34f21e55f683abdded0aa71df3958a49c8b69": {
			Hash: "dcc98e38e1edb55a97265efca6c34f21e55f683abdded0aa71df3958a49c8b69",
			URL:  "http://lightchain-sirius.prod.svc.cluster.local:8545",
		},
		"26a2800156f76b66bcb5661f2988a9d09e76caaffd053fe17bf20d251b4cb823": {
			Hash: "26a2800156f76b66bcb5661f2988a9d09e76caaffd053fe17bf20d251b4cb823",
			URL:  "http://lightchain-mainnet.prod.svc.cluster.local:8545",
		},
		"73d8dd1b7d8aa02254e75936b09780447c06729f3e55f7ae5eb94ab732c1ec05": {
			Hash: "73d8dd1b7d8aa02254e75936b09780447c06729f3e55f7ae5eb94ab732c1ec05",
			URL:  "http://particl-mainnet.prod.svc.cluster.local:51735",
		},
		"6cbb58da0b05d23022557dd2e479dd5cdf2441f20507b37383467d837ad40f5e": {
			Hash: "6cbb58da0b05d23022557dd2e479dd5cdf2441f20507b37383467d837ad40f5e",
			URL:  "http://pivx-mainnet.prod.svc.cluster.local:51473",
		},
		"54cb0d71117aa644e74bdea848d61bd2fd410d3d4a3ed92b46b0847769dc132e": {
			Hash: "54cb0d71117aa644e74bdea848d61bd2fd410d3d4a3ed92b46b0847769dc132e",
			URL:  "http://poa-core-2.prod.svc.cluster.local:8545",
		},
		"cb92cb81d6f72f55114140a7bbe5e0f63524d1200fe63250f58dfe5d907032bf": {
			Hash: "cb92cb81d6f72f55114140a7bbe5e0f63524d1200fe63250f58dfe5d907032bf",
			URL:  "http://poa-kovan-2.prod.svc.cluster.local:8545",
		},
		"e458822c5f4d927c29aa4240a34647e11aff75232ccb9ffb50af06dc4469a5fa": {
			Hash: "e458822c5f4d927c29aa4240a34647e11aff75232ccb9ffb50af06dc4469a5fa",
			URL:  "http://poa-sokol-2.prod.svc.cluster.local:8545",
		},
		"0dfcabfb7f810f96cde01d65f775a565d3a60ad9e15575dfe3d188ff506c35a0": {
			Hash: "0dfcabfb7f810f96cde01d65f775a565d3a60ad9e15575dfe3d188ff506c35a0",
			URL:  "http://poa-xdai-2.prod.svc.cluster.local:8545",
		},
		"866d7183a24fad1d0a32c399cf2a1101f3a3bdfdff999e142bd8f49b2ebc45d4": {
			Hash: "866d7183a24fad1d0a32c399cf2a1101f3a3bdfdff999e142bd8f49b2ebc45d4",
			URL:  "http://stellar-mainnet.prod.svc.cluster.local",
		},
		"4c0437dda63eff39f85c60d62ac936045da5e610aca97a3793771e271578c534": {
			Hash: "4c0437dda63eff39f85c60d62ac936045da5e610aca97a3793771e271578c534",
			URL:  "http://tezos-alpohanet-2.prod.svc.cluster.local:8732",
		},
		"773eda9368243afe027062d771b08cebddf22e03451e0eb5ed0ff4460288847e": {
			Hash: "773eda9368243afe027062d771b08cebddf22e03451e0eb5ed0ff4460288847e",
			URL:  "http://tezos-mainnet-2.prod.svc.cluster.local:8732",
		},
		"d5ddbb1ca49249438f552dccfd01918ee1fbdc6457997a142c8cfd144b40cd15": {
			Hash: "d5ddbb1ca49249438f552dccfd01918ee1fbdc6457997a142c8cfd144b40cd15",
			URL:  "http://tomochain-mainnet.prod.svc.cluster.local:8545",
		},
		"4ecc78e62904c833ad5b727b9abf343a17d0d24fb27e9b5d2dd8c34361c23156": {
			Hash: "4ecc78e62904c833ad5b727b9abf343a17d0d24fb27e9b5d2dd8c34361c23156",
			URL:  "http://waves-mainnet.prod.svc.cluster.local:6870",
		},
		"d754973bdeab17eaed47729ee074ad87737c3ce51198263b8c4781568ea39e72": {
			Hash: "d754973bdeab17eaed47729ee074ad87737c3ce51198263b8c4781568ea39e72",
			URL:  "http://zcoin-mainnet.prod.svc.cluster.local:8888",
		},
	},
	}
)

func GenChains(homeDir string, hbs types.HostedBlockchains) string {
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

func DefaultTestnetChains() types.HostedBlockchains {
	return defaultChains
}
