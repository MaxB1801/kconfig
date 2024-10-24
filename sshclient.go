package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
)

func sshclient(username, password, ip, path string) []byte {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf(ip+":22"), config)
	if err != nil {
		log.Fatal("Failed to connect to ", ip, err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	r, err := session.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	// Define the command to read the remote file
	cmd := fmt.Sprintf("cat %s", path)

	if err := session.Start(cmd); err != nil {
		log.Fatal(err)
	}

	// Read the output into a variable
	content, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Failed to read output: %s", err)
	}

	if err := session.Wait(); err != nil {
		log.Fatal(err)
	}

	return content

}
