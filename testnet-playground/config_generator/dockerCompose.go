package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

const (
	OpenPocketRPCPort         = "8081:8081"
	OpenTendermintRPCPort     = "26657:26657"
	PocketRPCPort             = "8081"
	TendermintRPCPort         = "26657"
	TendermintPeersPort       = "26656"
	TendermintPeersPortLegacy = "46656"
	DockerVersion             = "2.1"
	Driver                    = "bridge"
	Image                     = `poktnetwork/pocket-core:${ENV:-staging-latest}`
	Privileged                = true
	Network                   = "pocket"
	Command                   = "/usr/bin/expect /home/app/.pocket/local_command.sh"
	Build                     = `../../docker`
	Volume1                   = "./local_command.sh:/home/app/.pocket/local_command.sh"
	Volume2                   = "./chains.json:/home/app/.pocket/config/chains.json"
	Volume3                   = "./genesis.json:/home/app/.pocket/config/genesis.json"
)

var (
	ports = []string{
		OpenPocketRPCPort,
		OpenTendermintRPCPort,
	}
	expose = []string{
		PocketRPCPort,
		TendermintRPCPort,
		TendermintPeersPort,
		TendermintPeersPortLegacy,
	}
)

type DockerComposeFile struct {
	Version  string             `yaml:"version"`
	Services map[string]Testnet `yaml:"services"`
	Networks PocketNetworkWrap  `yaml:"networks"`
}

type PocketNetworkWrap struct {
	PocketNetwork PocketNetwork `yaml:"pocket"`
}

type PocketNetwork struct {
	Driver string `yaml:"driver"`
}

//
//type TestnetService struct {
//	Testnet `yaml:"pocket-core-testnet"`
//}

type Testnet struct {
	Image       string      `yaml:"image"`
	Privileged  bool        `yaml:"privileged"`
	Command     string      `yaml:"command"`
	Build       string      `yaml:"build"`
	Ports       []string    `yaml:"ports"`
	Expose      []string    `yaml:"expose"`
	Environment KeyAndSeeds `yaml:"environment"`
	Networks    []string    `yaml:"networks"`
	Volumes     []string    `yaml:"volumes"`
}

type KeyAndSeeds struct {
	Key   string `yaml:"POCKET_CORE_KEY"`
	Seeds string `yaml:"POCKET_CORE_SEEDS"`
}

func generateDockerComposeFile(homeDir string, keys KeysFile) {
	dockerComposeFile := DockerComposeFile{
		Version:  DockerVersion,
		Services: make(map[string]Testnet, 0),
		Networks: PocketNetworkWrap{
			PocketNetwork: PocketNetwork{
				Driver: Driver,
			},
		},
	}
	testnetServiceTemplate := Testnet{
		Image:      Image,
		Privileged: Privileged,
		Ports:      ports,
		Command:    Command,
		Build:      Build,
		Expose:     expose,
		Networks:   []string{Network},
		Volumes:    []string{Volume1, Volume2, Volume3},
	}
	for i, nodeKey := range keys.NodeKeys {
		// edit structure tag
		tn := testnetServiceTemplate
		if i != 0 {
			tn.Ports = []string{}
		}
		var seedString string
		for j, nk := range keys.NodeKeys {
			if nk.Addr == nodeKey.Addr {
				continue
			}
			seedString = seedString + strings.ToLower(nk.Addr) + fmt.Sprintf("@pocket-core-testnet%d", j) + ":" + TendermintPeersPort+" , "
		}
		ks := KeyAndSeeds{
			Key:   nodeKey.Priv,
			Seeds: seedString,
		}
		tn.Environment = ks
		dockerComposeFile.Services[fmt.Sprintf("pocket-core-testnet%d", i)] = tn
	}
	bz, err := yaml.Marshal(dockerComposeFile)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(homeDir+fs+"docker-compose.yaml", bz, os.ModePerm)
	if err != nil {
		fmt.Println(DirectoryCreationError.Error() + err.Error())
		os.Exit(0)
	}
}

var counter = 0

type TempTestnetService struct {
	Testnet `yaml:"pocket-core-testnet"`
}