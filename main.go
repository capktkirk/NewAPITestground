package main

import (
	"fmt"
	"os"
	"sync"
)

type KeyValueDB struct {
	Entries map[string]Payload `json:"entries"`
	mu      sync.Mutex
}
type Payload struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("You cannot run this program without a command and a key")
	}

	err := checkArgs(os.Args[1])
	if err != nil {
		fmt.Printf("Error with arguments: %s\n", err)
	}
}

func checkArgs(cmd string) error {
	valid := []string{
		"set",
		"get",
		"del",
		"ts",
	}
	if !slices.Contains(valid, cmd) {
		return fmt.Errorf("%s is not a valid command, use set, get, del or ts instead.", cmd)
	}
	return nil
}
