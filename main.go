package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] != "-edit" {
			fmt.Println("Use -edit to edit the kconfig.yaml")
			return
		}
		openEditor("kconfig.yaml")
		return
	}

	// Get the executable's directory
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err) // Handle the error if getting the executable path fails
	}

	// Get the directory of the executable
	exeDir := filepath.Dir(exePath)

	filedata, err := os.ReadFile(filepath.Join(exeDir, "kconfig.yaml"))
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
		configArray = append(configArray, sshclient(cluster.Username, cluster.Password, cluster.IP, cluster.Type))
	}

	saveData := formatKConfig(configArray, kconfig)

	for _, path := range kconfig.Savedir.Filepaths {
		createFile(saveData, filepath.Join(path, "config"))
	}
}
