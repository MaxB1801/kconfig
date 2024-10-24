package main

import (
	"fmt"
	"log"
	"os"
)

func createFile(config Config, filepath string) {
	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	n, err := f.WriteString(fmt.Sprintf("\n%s", config.Clusters[0].ServerContext.Server))
	if err != nil {
		log.Fatal("Error Appending to File", n, err)
	}

	f.Close()
}
