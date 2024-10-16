package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slices"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
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

	rootCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	start(rootCtx)
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

func start(ctx context.Context) {
	r := mux.NewRouter()
	r.HandleFunc("/", NewHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func NewHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
