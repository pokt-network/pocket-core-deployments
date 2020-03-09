package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/pokt-network/pocket-core-deployments/testnet-playground/config_generator/app"
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	WelcomeMessage = `
Welcome to the testnet playground generator!
Before I can generate your testnet, I need to ask you a few questions...`
	ModePrompt = `
Which mode would you like to run the generator in (1) custom local (2) for testnet (pocket inc. only)`
	NumberOfNodesPrompt = `
How many Nodes (validators) would you like pre-staked in this network?`
	NumberOfAppsPrompt = `
How many Applications (DApps) would you like pre-staked in this network?`
	NumberOfAccountsPrompt = `
How many accounts (unstaked actors) would you like in this network?
NOTE: all nodes and apps are already given an amount of unstaked coins too!`
	NewChainPrompt = `
Would you like to enter a non-native chain? (1) yes (2) no`
	ChainHashPrompt = `
Enter the chain hash:`
	URLPrompt = `
Enter the url of the chain:`
)

var (
	ReadInError = errors.New(`Uh oh, an error occurred reading in the information: `)
	fs          = string(filepath.Separator)
)

func main() {
	setup(gatherParameters())
}

func setup(mode, numberOfNodes, numberOfApps, numberOfAccounts int, hc types.HostedBlockchains, keys app.KeysFile) {
	home := generateTestnetHome()
	if numberOfNodes != 0 {
		keys = app.GenKeys(home, numberOfNodes, numberOfApps, numberOfAccounts)
	}
	chains := app.GenChains(home, hc)
	genesis := app.GenGenesis(home, keys, hc)
	app.GenDockerConfig(home, keys)
	app.WriteLocalCmd(home)
	app.NewKubenetesFile(home, keys, chains, genesis)
	app.GenFinishedMessages(keys)
}

func gatherParameters() (mode, numberOfNodes, numberOfApps, numberOfAccounts int, chains types.HostedBlockchains, keys app.KeysFile) {
	fmt.Println(WelcomeMessage)
	fmt.Println(ModePrompt)
	_, err := fmt.Scanf("%d", &mode)
	if err != nil {
		fmt.Println(ReadInError.Error() + err.Error())
		os.Exit(1)
	}
	keys = app.GetKeyFromFile()
	if len(keys.NodeKeys) == 0 {
		fmt.Println(NumberOfNodesPrompt)
		_, err = fmt.Scanf("%d", &numberOfNodes)
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
	}
	chains.M = make(map[string]types.HostedBlockchain)
	if mode == 2 {
		chains = app.DefaultTestnetChains()
	} else {
		var newChain int
		for {
			fmt.Println(NewChainPrompt)
			_, err = fmt.Scanf("%d", &newChain)
			if err != nil {
				fmt.Println(ReadInError.Error() + err.Error())
				os.Exit(1)
			}
			if newChain == 2 {
				break
			}
			reader := bufio.NewReader(os.Stdin)
			fmt.Println(ChainHashPrompt)
			hash, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Println(URLPrompt)
			url, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			chains.M[strings.Trim(hash, "\n")] = types.HostedBlockchain{
				Hash: strings.Trim(hash, "\n"),
				URL:  strings.Trim(url, "\n"),
			}
		}
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
