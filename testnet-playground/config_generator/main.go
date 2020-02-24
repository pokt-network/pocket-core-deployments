package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/pokt-network/pocket-core-deployments/testnet-playground/config_generator/app"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	WelcomMessage = `
Welcome to the testnet playground generator!

Befofe I can generate your testnet, I need to ask you a few questions...`
	NumberOfNodesPrompt = `
How many Nodes (validators) would you like pre-staked in this network?`
	NumberOfAppsPrompt = `
How many Applications (DApps) would you like pre-staked in this network?`
	NumberOfAccountsPrompt = `
How many accounts (unstaked actors) would you like in this network?
NOTE: all nodes and apps are already given an amount of unstaked coins too!`
	URLForEthereumNodePrompt = `
For testing purposes, we need you to point us to an ethereum node!
NOTE: if you don't have an ethereum node, just put an arbitrary url! 
Be sure to not do any relay requests to ethereum then!'`
	URLForBitcoinNodePrompt = `
For testing purposes, we need you to point us to a bitcoin node!
NOTE: if you don't have a bitcoin node, just put an arbitrary url! 
Be sure to not do any relay requests to bitcoin then!'`
	MinutesTillGenesisPrompt = `
And finally, how many minutes until genesis starts?
NOTE: this is important! Nodes will start mining blocks at this exact time!
If all genesis validators are not up, the network will never start`
)

var (
	ReadInError            = errors.New(`Uh oh, an error occurred reading in the information: `)
	fs                     = string(filepath.Separator)
)

func main() {
	setup(gatherParameters())
}

func setup(numberOfNodes, numberOfApps, numberOfAccounts, minutesTillGenesisStart int, ethereumURL, bitcoinURL string) {
	home := generateTestnetHome()
	keys := app.GenKeys(home, numberOfNodes, numberOfApps, numberOfAccounts)
	genesis:= app.GenGenesis(home, keys, minutesTillGenesisStart)
	chains := app.GenChains(home, ethereumURL, bitcoinURL)
	app.GenDockerConfig(home, keys)
	app.WriteLocalCmd(home)
	app.NewKubenetesFile(home, keys, chains, genesis)
	app.GenFinishedMessages(keys)
}

func gatherParameters() (numberOfNodes, numberOfApps, numberOfAccounts, minutesTillGenesisStart int, ethereumURL, bitcoinURL string) {
	fmt.Println(WelcomMessage)
	fmt.Println(NumberOfNodesPrompt)
	_, err := fmt.Scanf("%d", &numberOfNodes)
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	fmt.Println(NumberOfAppsPrompt)
	_, err = fmt.Scanf("%d", &numberOfApps)
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	fmt.Println(NumberOfAccountsPrompt)
	_, err = fmt.Scanf("%d", &numberOfAccounts)
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	fmt.Println(URLForEthereumNodePrompt)
	reader := bufio.NewReader(os.Stdin)
	ethereumURL, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	fmt.Println(URLForBitcoinNodePrompt)
	bitcoinURL, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	fmt.Println(MinutesTillGenesisPrompt)
	_, err = fmt.Scanf("%d", &minutesTillGenesisStart)
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	return
}

func generateTestnetHome() (directoryName string) {
	directoryName = "docker-compose" + fs + "testnet-playground_" + strconv.Itoa(time.Now().Nanosecond())
	err := os.Mkdir(directoryName, os.ModePerm)
	if err != nil {
		fmt.Println(app.DirectoryCreationError.Error() + err.Error())
		os.Exit(1)
	}
	return
}
