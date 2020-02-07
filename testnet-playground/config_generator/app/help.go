package app

import "fmt"

const (
	FinishedMessage = `Congratulations! Your network cluster is created!

How to run?
  - cd ../../docker-compose/testnet-playground_ <TAB>
  - docker-compose build
  - docker-compose up`
	ImportAccountMessage = `
In your pocket core CLI instance, run ./pocket-core accounts import-raw `
	FutureMessage        = `
Now test commands through CLI and RPC <localhost:8081>`
)

func GenFinishedMessages(keys KeysFile) {
	fmt.Println(FinishedMessage)
	fmt.Println(ImportAccountMessage + keys.NodeKeys[0].Priv)
	fmt.Println(FutureMessage)
}
