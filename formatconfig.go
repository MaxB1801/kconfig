package main

import (
	"log"
	"regexp"

	"gopkg.in/yaml.v2"
)

func formatKConfig(configArray [][]byte, kconfig KConfig) []byte {

	// Initialize the slice to hold the unmarshaled Configs
	ServerConfigs := make([]Config, len(configArray))
	var err error
	for n, config := range configArray {
		err = yaml.Unmarshal(config, &ServerConfigs[n])
		if err != nil {
			log.Fatal("Error Parsing SSH Config")
		}
	}

	// combine the configs and the marshal to file using the yaml package
	var finalConfig Config
	// Define a regex pattern to match the IP part of the URL
	re := regexp.MustCompile(`(?m)://[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+`)
	finalConfig.CurrentContext = kconfig.Clusters[0].Name
	finalConfig.Kind = "Config"
	for n, config := range ServerConfigs {
		config.Clusters[0].Name = kconfig.Clusters[n].Name
		config.Clusters[0].ServerContext.Server = re.ReplaceAllString(config.Clusters[0].ServerContext.Server, "://"+kconfig.Clusters[n].IP)
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
		log.Fatal("Error Reading Newly Made KubeConfig ", err)
	}

	return saveData

}
