package app

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	ApiVersion               = "v1"
	ServiceKind              = "Service"
	KomposeCMD               = "kompose convert -f docker-compose.yml -o geth.yml"
	KomposeVersion           = "1.18.0 (06a2e56)"
	PocketCoreTestnet        = "pocket-core-testnet"
	HTTP                     = "http"
	TCP                      = "TCP"
	KubernetesKind           = "List"
	Peers                    = "peers"
	LegacyPeersPort          = "46656"
	Tendermint               = "tendermint"
	TendermintPeers          = "tendermint-peers"
	LoadBalancer             = "LoadBalancer"
	LoadbalancerIP           = "35.236.208.175"
	APIAppsVersion           = "apps/v1"
	DeploymentKind           = "Deployment"
	Replicas                 = "1"
	ConfigMapResourceVersion = "1"
	ConfigMapKind            = "ConfigMap"
	TypeRecreate             = "Recreate"
	FsGroup                  = "100"
	PocketCoreStart          = "pocket-core start"
	PocketCoreSeedsFlag      = "--seeds"
	PocketCoreSeeds          = "$(POCKET_CORE_SEEDS)"
	BlockTimeFlag            = "--blockTime=1"
	PCImage                  = "poktnetwork/pocket-core:staging-latest"
	ImagePolicyRestart       = "Always"
	PullPolicy               = "Always"
	PCKey                    = "POCKET_CORE_KEY"
	PCPassPhrase             = "POCKET_CORE_PASSPHRASE"
	PCPassphraseKey          = "pocket_core_passphrase"
	Chains                   = "POCKET_CORE_CHAINS"
	PCChainsKey              = "pocket_core_chains"
	PCGenesis                = "POCKET_CORE_GENESIS"
	PCGenesisKey             = "pocket_core_genesis"
	PCSeeds                  = "POCKET_CORE_SEEDS"
	PCSeedsKey               = "pocket_core_seeds"
	CPULimit                 = "500m"
	MemoryLimit              = "500Mi"
	CPURequests              = "250m"
	MemoryRequests           = "300Mi"
	TerminalGracePeriod      = "90"
	PocketCore               = "pocket-core"
	PCPass                   = "VVb4HgtVQGOKhj7UOnJq"
)

var ServicePorts = []Port{
	{
		Name:       HTTP,
		Protocol:   TCP,
		Port:       PocketRPCPort,
		TargetPort: PocketRPCPort,
	},
	{
		Name:       Peers,
		Protocol:   TCP,
		Port:       LegacyPeersPort,
		TargetPort: LegacyPeersPort,
	},
	{
		Name:       Tendermint,
		Protocol:   TCP,
		Port:       TendermintRPCPort,
		TargetPort: TendermintRPCPort,
	},
	{
		Name:       TendermintPeers,
		Protocol:   TCP,
		Port:       TendermintPeersPort,
		TargetPort: TendermintPeersPort,
	},
}

var Args = []string{
	PocketCoreStart,
	PocketCoreSeedsFlag,
	PocketCoreSeeds,
	BlockTimeFlag,
}

func NewKubenetesFile(homeDir string, keys KeysFile, chainsJSON string, genesisJSON string) {
	var items []Item
	var seeds string
	for i, key := range keys.NodeKeys {
		pcTestnet := PocketCoreTestnet
		if i != 0 {
			pcTestnet += strconv.Itoa(i)
		}
		items = append(items, ServiceItem{
			ApiVersion: ApiVersion,
			Kind:       ServiceKind,
			Metadata: Metadata{
				Annotations: Annotations{
					KomposeCmd:     KomposeCMD,
					KomposeVersion: KomposeVersion,
				},
				CreationTimestamp: nil,
				TemplateLabels: TemplateLabels{
					IOKomposeService: pcTestnet,
					App:              pcTestnet,
				},
				Name: pcTestnet,
			},
			Spec: Spec{
				Ports:          ServicePorts,
				Type:           LoadBalancer,
				LoadBalancerIP: LoadbalancerIP,
				Selector: Selector{
					IOKomposeService: pcTestnet,
				},
			},
			Status: Status{
				LoadBalancer: struct{}{},
			},
		})
		var envs = []Env{
			EnvValueFrom{
				Name: Chains,
				ValueFrom: ValueFrom{ConfigMapKeyRef{
					Name: PocketCoreTestnet,
					Key:  PCChainsKey,
				}},
			},
			EnvValueFrom{
				Name: PCPassPhrase,
				ValueFrom: ValueFrom{ConfigMapKeyRef{
					Name: PocketCoreTestnet,
					Key:  PCPassphraseKey,
				}},
			},
			EnvValueFrom{
				Name: PCSeeds,
				ValueFrom: ValueFrom{ConfigMapKeyRef{
					Name: PocketCoreTestnet,
					Key:  PCSeedsKey,
				}},
			},
			EnvValueFrom{
				Name: PCGenesis,
				ValueFrom: ValueFrom{ConfigMapKeyRef{
					Name: PocketCoreTestnet,
					Key:  PCGenesisKey,
				}},
			},
			EnvValue{
				Name:  PCKey,
				Value: key.Priv,
			},
		}
		items = append(items, DeploymentItem{
			ApiVersion: APIAppsVersion,
			Kind:       DeploymentKind,
			Metadata: Metadata{
				Annotations: Annotations{
					KomposeCmd:     KomposeCMD,
					KomposeVersion: KomposeVersion,
				},
				CreationTimestamp: nil,
				TemplateLabels: TemplateLabels{
					IOKomposeService: pcTestnet,
					App:              pcTestnet,
				},
				Name: pcTestnet,
			},
			DeploymentSpec: DeploymentSpec{
				Replicas: Replicas,
				Selector: DeploymentSelector{
					MatchLabels: MatchLabels{
						App: pcTestnet,
					},
				},
				Strategy: Strategy{
					Type: TypeRecreate,
				},
				Template: Template{
					TemplateMetadata: TemplateMetadata{
						TemplateLabels: TemplateLabels{
							App:              pcTestnet,
							IOKomposeService: pcTestnet,
						},
					},
					Spec: TemplateSpec{
						SecurityContext: SecurityContext{
							FsGroup: FsGroup,
						},
						Containers: Containers{
							Name:            pcTestnet,
							Args:            Args,
							Image:           PCImage,
							ImagePullPolicy: PullPolicy,
							Envs:            envs,
							Ports: []ContainerPorts{
								{
									ContainerPort: PocketRPCPort,
								},
								{
									ContainerPort: TendermintPeersPortLegacy,
								},
								{
									ContainerPort: TendermintPeersPort,
								},
								{
									ContainerPort: TendermintRPCPort,
								},
							},
							Resources: Resources{
								Limits: CPUMem{
									CPU:    CPULimit,
									Memory: MemoryLimit,
								},
								Requests: CPUMem{
									CPU:    CPURequests,
									Memory: MemoryRequests,
								},
							},
							RestartPolicy: ImagePolicyRestart,
							NodeSelector: NodeSelector{
								CloudGoogle: PocketCore,
							},
							TerminationGracePeriodSeconds: TerminalGracePeriod,
						},
					},
				},
			},
		})
		seeds = seeds + key.Pub + "@" + pcTestnet + ":" + TendermintPeersPort + ", "
	}

	items = append(items, ConfigMapItem{
		ApiVersion: ApiVersion,
		Kind:       ConfigMapKind,
		ConfigMapMetaData: ConfigMapMetaData{
			Name:            PocketCoreTestnet,
			ResourceVersion: ConfigMapResourceVersion,
		},
		Data: Data{
			PocketCoreSeeds:      seeds,
			PocketCorePassphrase: PCPass,
			PocketCoreChains:     chainsJSON,
			PocketCoreGenesis:    genesisJSON,
		},
	})
	file := Kubernetes{
		ApiVersion: ApiVersion,
		Items:      items,
		Kind:       KubernetesKind,
		Metadata:   struct{}{},
	}
	bz, err := yaml.Marshal(file)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(homeDir+fs+"pocket-core-testnet.yaml", bz, os.ModePerm)
	if err != nil {
		fmt.Println(DirectoryCreationError.Error() + err.Error())
		os.Exit(0)
	}
}

type Kubernetes struct {
	ApiVersion string   `yaml:"apiVersion"`
	Items      []Item   `yaml:"items"`
	Kind       string   `yaml:"kind"`
	Metadata   struct{} `yaml:"metadata"`
}

var _ Item = ServiceItem{}
var _ Item = ConfigMapItem{}
var _ Item = DeploymentItem{}

type Item interface{}

type ServiceItem struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   `yaml:"metadata"`
	Spec       `yaml:"spec"`
	Status     `yaml:"status"`
}

type Metadata struct {
	Annotations       `yaml:"annotations"`
	CreationTimestamp *time.Time `yaml:"creationTimestamp"`
	TemplateLabels    `yaml:"labels"`
	Name              string `yaml:"name"`
}

type Spec struct {
	Ports          []Port `yaml:"ports"`
	Type           string `yaml:"type"`
	LoadBalancerIP string `yaml:"loadBalancerIP"`
	Selector       `yaml:"selector"`
}

// config map item

type ConfigMapItem struct {
	ApiVersion        string `yaml:"apiVersion"`
	Kind              string `yaml:"kind"`
	ConfigMapMetaData `yaml:"metadata"`
	Data              `yaml:"data"`
}

type Data struct {
	PocketCoreSeeds      string `yaml:"pocket_core_seeds"`
	PocketCorePassphrase string `yaml:"pocket_core_passphrase"`
	PocketCoreChains     string `yaml:"pocket_core_chains"`
	PocketCoreGenesis    string `yaml:"pocket_core_genesis"`
}

type ConfigMapMetaData struct {
	Name            string `yaml:"name"`
	ResourceVersion string `yaml:"resourceVersion"`
}

// deployment item

type DeploymentItem struct {
	ApiVersion     string `yaml:"apiVersion"`
	Kind           string `yaml:"kind"`
	Metadata       `yaml:"metadata"`
	DeploymentSpec `yaml:"spec"`
}

type DeploymentSpec struct {
	Replicas string             `yaml:"replicas"`
	Selector DeploymentSelector `yaml:"selector"`
	Strategy `yaml:"strategy"`
	Template `yaml:"template"`
}

type Template struct {
	TemplateMetadata `yaml:"metadata"`
	Spec             TemplateSpec `yaml:"spec"`
}

type TemplateSpec struct {
	SecurityContext `yaml:"securityContext"`
	Containers      `yaml:"containers"`
}

type Containers struct {
	Name                          string           `yaml:"name"`
	Args                          []string         `yaml:"args"`
	Image                         string           `yaml:"image"`
	ImagePullPolicy               string           `yaml:"imagePullPolicy"`
	Envs                          []Env            `yaml:"env"`
	Ports                         []ContainerPorts `yaml:"ports"`
	Resources                     `yaml:"resources"`
	RestartPolicy                 string `yaml:"restartPolicy"`
	NodeSelector                  `yaml:"nodeSelector"`
	TerminationGracePeriodSeconds string `yaml:"terminationGracePeriodSeconds"`
}

type NodeSelector struct {
	CloudGoogle string `yaml:"cloud.google.com/gke-nodepool"`
}

type Resources struct {
	Limits   CPUMem `yaml:"limits"`
	Requests CPUMem `yaml:"requests"`
}

type CPUMem struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type ContainerPorts struct {
	ContainerPort string `yaml:"containerPort"`
}

type Env interface {
}

var _ Env = EnvValue{}
var _ Env = EnvValueFrom{}

type EnvValueFrom struct {
	Name      string `yaml:"name"`
	ValueFrom `yaml:"valueFrom"`
}

type EnvValue struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type ValueFrom struct {
	ConfigMapKeyRef `yaml:"configMapKeyRef"`
}

type ConfigMapKeyRef struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

type SecurityContext struct {
	FsGroup string `yaml:"fsGroup"`
}

type TemplateMetadata struct {
	TemplateLabels    `yaml:"labels"`
	CreationTimestamp *time.Time `yaml:"creationTimestamp"`
}

type TemplateLabels struct {
	App              string `yaml:"app"`
	IOKomposeService string `yaml:"io.kompose.service"`
}

type Strategy struct {
	Type string `yaml:"type"`
}

type DeploymentSelector struct {
	MatchLabels `yaml:"matchLabels"`
}

type MatchLabels struct {
	App string `yaml:"app"`
}

type Status struct {
	LoadBalancer struct{} `yaml:"loadBalancer"`
}

type Annotations struct {
	KomposeCmd     string `yaml:"kompose.cmd"`
	KomposeVersion string `yaml:"kompose.version"`
}

type Port struct {
	Name       string `yaml:"name"`
	Protocol   string `yaml:"protocol"`
	Port       string `yaml:"port"`
	TargetPort string `yaml:"targetPort"`
}

type Selector struct {
	IOKomposeService string `yaml:"io.kompose.service"`
}
