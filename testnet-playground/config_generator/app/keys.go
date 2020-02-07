package app

import (
	"encoding/json"
	"fmt"
	"github.com/pokt-network/posmint/crypto"
	"io/ioutil"
	"os"
)

type JSONKeys struct {
	Priv string `json:"private_key"`
	Pub  string `json:"public_key"`
	Addr string `json:"address"`
}

type KeysFile struct {
	NodeKeys []JSONKeys `json:"nodes"`
	AppKeys  []JSONKeys `json:"apps"`
	AccKeys  []JSONKeys `json:"accs"`
}

func GenKeys(directoryName string, numberOfNodes, numberOfApps, numberOfAccounts int) KeysFile {
	res := KeysFile{
		NodeKeys: make([]JSONKeys, 0),
		AppKeys:  make([]JSONKeys, 0),
		AccKeys:  make([]JSONKeys, 0),
	}
	for i := 0; i < numberOfNodes; i++ {
		pk := crypto.GenerateEd25519PrivKey()
		res.NodeKeys = append(res.NodeKeys, JSONKeys{
			Priv: pk.RawString(),
			Pub:  pk.PublicKey().RawString(),
			Addr: pk.PublicKey().Address().String(),
		})
	}
	for i := 0; i < numberOfApps; i++ {
		pk := crypto.GenerateEd25519PrivKey()
		res.AppKeys = append(res.AppKeys, JSONKeys{
			Priv: pk.RawString(),
			Pub:  pk.PublicKey().RawString(),
			Addr: pk.PublicKey().Address().String(),
		})
	}
	for i := 0; i < numberOfAccounts; i++ {
		pk := crypto.GenerateEd25519PrivKey()
		res.AccKeys = append(res.AccKeys, JSONKeys{
			Priv: pk.RawString(),
			Pub:  pk.PublicKey().RawString(),
			Addr: pk.PublicKey().Address().String(),
		})
	}
	bz, _ := json.MarshalIndent(res, "", "  ")
	err := ioutil.WriteFile(directoryName+fs+"keys.json", bz, os.ModePerm)
	if err != nil {
		fmt.Println(DirectoryCreationError.Error() + err.Error())
		os.Exit(0)
	}
	return res
}
