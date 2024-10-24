package main

import (
	"fmt"
	"log"
	"os"
)

func createFile(config []byte, filepath string) {
	// Delete the file if it exists
	if err := os.Remove(filepath); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error deleting file: %v", err)
	}

	// Create a new file
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write(config)
	if err != nil {
		fmt.Printf("error writing to file: %v\n", err)
		return
	}

}
