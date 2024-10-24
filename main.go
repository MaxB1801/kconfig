package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// define structs for input yaml
type Cluster struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	IP       string `yaml:"ip"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type KConfig struct {
	Savedir struct {
		Filepaths []string `yaml:"filepaths"`
	} `yaml:"savedir"`
	Clusters []Cluster `yaml:"clusters"`
}

// define struct for the ssh config
type Config struct {
	Clusters       []Clusters  `yaml:"clusters"`
	Contexts       []Context   `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
	Kind           string      `yaml:"kind"`
	Preferences    Preferences `yaml:"preferences"`
	Users          []User      `yaml:"users"`
}

type Clusters struct {
	Name          string `yaml:"name"`
	ServerContext struct {
		CertificateAuthorityData string `yaml:"certificate-authority-data"`
		Server                   string `yaml:"server"`
	} `yaml:"cluster"`
}

type Context struct {
	Name    string `yaml:"name"`
	Context struct {
		Cluster string `yaml:"cluster"`
		User    string `yaml:"user"`
	} `yaml:"context"`
}

type Preferences struct{}

type User struct {
	Name string `yaml:"name"`
	User struct {
		ClientCertificateData string `yaml:"client-certificate-data"`
		ClientKeyData         string `yaml:"client-key-data"`
	} `yaml:"user"`
}

func main() {
	filedata, err := os.ReadFile("kconfig.yaml")
	if err != nil {
		log.Fatal("Error Reading kconfig.yaml")
	}

	var kconfig KConfig
	err = yaml.Unmarshal(filedata, &kconfig)
	if err != nil {
		log.Fatal("Error Parsing kconfig.yaml")
	}

	var configArray [][]byte
	for _, cluster := range kconfig.Clusters {
		configArray = append(configArray, sshclient(cluster.Username, cluster.Password, cluster.IP, "/etc/rke2/rke2.yaml"))
	}

	// Initialize the slice to hold the unmarshaled Configs
	ServerConfigs := make([]Config, len(configArray))
	for n, config := range configArray {
		err = yaml.Unmarshal(config, &ServerConfigs[n])
		if err != nil {
			log.Fatal("Error Parsing SSH Config")
		}
	}

	// combine the configs and the marshal to file using the yaml package
	var finalConfig Config
	finalConfig.CurrentContext = kconfig.Clusters[0].Name
	finalConfig.CurrentContext = kconfig.Clusters[0].Name
	for n, config := range ServerConfigs {
		config.Clusters[0].Name = kconfig.Clusters[n].Name
		// config.Clusters[0].clusters.f
		finalConfig.Clusters = append(finalConfig.Clusters, config.Clusters[0])
		config.Contexts[0].Context.Cluster = kconfig.Clusters[n].Name
		config.Contexts[0].Context.User = kconfig.Clusters[n].Name
		config.Contexts[0].Name = kconfig.Clusters[n].Name
		finalConfig.Contexts = append(finalConfig.Contexts, config.Contexts[0])
		config.Users[0].Name = kconfig.Clusters[n].Name
		finalConfig.Users = append(finalConfig.Users, config.Users[0])
	}

	// Marshal the config to YAML
	saveData, err := yaml.Marshal(finalConfig)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	createFile(saveData, filepath.Join(kconfig.Savedir.Filepaths[0], "config"))

}
